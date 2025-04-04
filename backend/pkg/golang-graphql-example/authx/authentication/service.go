package authentication

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/models"
	cerrors "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/samber/lo"
	"golang.org/x/oauth2"
)

const (
	callbackPath           = "/auth/oidc/callback"
	loginPath              = "/auth/oidc"
	logoutPath             = "/auth/oidc/logout"
	userContextKeyName     = "USER_CONTEXT_KEY"
	redirectQueryKey       = "rd"
	stateRedirectSeparator = ":"
	stateLength            = 2
)

var userContextKey = &contextKey{name: userContextKeyName}

type providerEndpointsClaims struct {
	EndSessionEndpointURL *url.URL
	EndSessionEndpoint    string `json:"end_session_endpoint"`
}

type service struct {
	verifier   *oidc.IDTokenVerifier
	cfgManager config.Manager
}

// GetAuthenticatedUser will get authenticated user in context.
func GetAuthenticatedUserFromContext(ctx context.Context) *models.OIDCUser {
	res, _ := ctx.Value(userContextKey).(*models.OIDCUser)

	return res
}

// GetAuthenticatedUser will get authenticated user in context.
func GetAuthenticatedUserFromGin(c *gin.Context) *models.OIDCUser {
	res, _ := c.Get(userContextKeyName)
	res1, _ := res.(*models.OIDCUser)

	return res1
}

// SetAuthenticatedUserToContext will set user in context.
func SetAuthenticatedUserToContext(ctx context.Context, us *models.OIDCUser) context.Context {
	return context.WithValue(ctx, userContextKey, us)
}

// SetAuthenticatedUserToGin will set user in gin context.
func SetAuthenticatedUserToGin(c *gin.Context, us *models.OIDCUser) {
	c.Set(userContextKeyName, us)
}

// OIDCEndpoints will set OpenID Connect endpoints for authentication and callback.
func (s *service) OIDCEndpoints(router gin.IRouter) error {
	ctx := context.Background()

	// Get configuration
	cfg := s.cfgManager.GetConfig()

	// Create provider
	provider, err := oidc.NewProvider(ctx, cfg.OIDCAuthentication.IssuerURL)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Create provider endpoints claims
	pec := &providerEndpointsClaims{}
	// Get claims
	err = provider.Claims(pec)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Parse logout endpoint
	eseURL, err := url.Parse(pec.EndSessionEndpoint)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Save url
	pec.EndSessionEndpointURL = eseURL

	oidcConfig := &oidc.Config{
		ClientID: cfg.OIDCAuthentication.ClientID,
	}
	// Create verifier
	verifier := provider.Verifier(oidcConfig)

	// Build redirect url
	mainRedirectURLObject, err := url.Parse(cfg.OIDCAuthentication.RedirectURL)
	// Check if error exists
	if err != nil {
		return errors.WithStack(err)
	}
	// Continue to build redirect url
	mainRedirectURLObject.Path = path.Join(mainRedirectURLObject.Path, callbackPath)
	mainRedirectURLStr := mainRedirectURLObject.String()

	// Create OIDC configuration
	oauthConfig := oauth2.Config{
		ClientID:    cfg.OIDCAuthentication.ClientID,
		Endpoint:    provider.Endpoint(),
		Scopes:      cfg.OIDCAuthentication.Scopes,
		RedirectURL: mainRedirectURLStr,
	}
	if cfg.OIDCAuthentication.ClientSecret != nil {
		oauthConfig.ClientSecret = cfg.OIDCAuthentication.ClientSecret.Value
	}

	// Store state
	state := cfg.OIDCAuthentication.State

	// Store provider verifier in map
	s.verifier = verifier

	// Login mount point
	router.GET(loginPath, func(c *gin.Context) {
		// Get redirect query from query params
		rdVal := c.Query(redirectQueryKey)
		// Build new state with redirect value
		// Same solution as here: https://github.com/oauth2-proxy/oauth2-proxy/blob/3fa42edb7350219d317c4bd47faf5da6192dc70f/oauthproxy.go#L751
		newState := state + stateRedirectSeparator + rdVal

		c.Redirect(http.StatusFound, oauthConfig.AuthCodeURL(newState))
		c.Abort()
	})

	// Logout mount point
	router.GET(logoutPath, func(c *gin.Context) {
		// Initialize redirect to
		rdTo := "/"
		// Check if logout url exists and logout redirect url exists
		if pec.EndSessionEndpoint != "" && cfg.OIDCAuthentication.LogoutRedirectURL != "" {
			// Parse logout url
			lgURL := *pec.EndSessionEndpointURL
			// Get params
			qs := lgURL.Query()
			// Add param
			qs.Add("redirect_uri", cfg.OIDCAuthentication.LogoutRedirectURL)
			// Save them
			lgURL.RawQuery = qs.Encode()
			// Encode
			rdTo = lgURL.String()
		}

		// Flush auth cookie
		flushAuthCookie(c, cfg)
		// Redirect
		c.Redirect(http.StatusFound, rdTo)
		c.Abort()
	})

	// Redirect mount point
	router.GET(mainRedirectURLObject.Path, func(c *gin.Context) {
		// Get logger from request
		logger := log.GetLoggerFromGin(c)

		// Get state from request
		reqQueryState := c.Query("state")
		// Check if state exists
		if reqQueryState == "" {
			err := cerrors.NewInvalidInputError("state not found in request")

			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		// Split request query state to get redirect url and original state
		split := strings.SplitN(reqQueryState, stateRedirectSeparator, stateLength)
		// Prepare and affect values
		reqState := split[0]
		rdVal := ""
		// Check if length is ok to include a redirect url
		if len(split) == stateLength {
			rdVal = split[1]
		}

		// Check state
		if reqState != state {
			err := cerrors.NewInvalidInputError("state did not match")
			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		// Check if rdVal exists and that redirect url value is valid
		if rdVal != "" {
			isValid, err := isValidRedirect(rdVal, utils.GetRequestURL(c.Request))
			// Check error
			if err != nil {
				// Answer
				logger.Error(err)
				utils.AnswerWithError(c, err)

				return
			}
			// Check if it is invalid
			if !isValid {
				err := cerrors.NewInvalidInputError(
					"redirect url is invalid",
					cerrors.WithPublicErrorMessage("redirect url is invalid"),
				)

				logger.Error(err)
				utils.AnswerWithError(c, err)

				return
			}
		}

		oauth2Token, err := oauthConfig.Exchange(ctx, c.Query("code"))
		if err != nil {
			err = cerrors.NewInternalServerError("failed to exchange token: " + err.Error())
			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			err = cerrors.NewInternalServerError("no id_token field in token")
			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			err = cerrors.NewInternalServerError("failed to verify ID Token: " + err.Error())
			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		var resp models.OIDCUser

		// Try to open JWT token in order to verify that we can open it
		err = idToken.Claims(&resp)
		if err != nil {
			// Create error with stack trace
			err2 := errors.WithStack(err)
			logger.Error(err2)
			utils.AnswerWithError(c, err2)

			return
		}

		resp.OriginalToken = rawIDToken
		// Now, we know that we can open jwt token to get claims

		// Build cookie
		cookie := &http.Cookie{
			Expires:  idToken.Expiry,
			Name:     cfg.OIDCAuthentication.CookieName,
			Value:    rawIDToken,
			HttpOnly: true,
			Secure:   cfg.OIDCAuthentication.CookieSecure,
			Path:     "/",
		}

		// Set cookie
		http.SetCookie(c.Writer, cookie)

		// Manage default redirect case
		if rdVal == "" {
			rdVal = "/"
		}

		logger.Info("Successful authentication detected")
		c.Redirect(http.StatusTemporaryRedirect, rdVal)
		c.Abort()
	})

	return nil
}

func (s *service) Middleware(unauthorizedPathRegexList []*regexp.Regexp) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get logger
		logger := log.GetLoggerFromGin(c)
		// Get configuration
		cfg := s.cfgManager.GetConfig()
		// Get JWT Token from header or cookie
		jwtContent, err := getJWTToken(logger, c.Request, cfg.OIDCAuthentication.CookieName)
		// Check if error exists
		if err != nil {
			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}
		// Check if JWT content is empty or not
		if jwtContent == "" {
			logger.Error(
				cerrors.NewUnauthorizedError(
					"No auth header or cookie detected, redirect to oidc login",
				),
			)
			redirectOrUnauthorized(c, unauthorizedPathRegexList)

			return
		}

		// Parse token

		var ouser models.OIDCUser
		// Verify token
		idToken, err := s.verifier.Verify(context.Background(), jwtContent)
		// Check error
		if err != nil {
			logger.Error(errors.WithStack(err))
			// Flush potential cookie
			flushAuthCookie(c, cfg)

			redirectOrUnauthorized(c, unauthorizedPathRegexList)

			return
		}

		// Get claims
		err = idToken.Claims(&ouser)
		if err != nil {
			logger.Error(errors.WithStack(err))
			// Flush potential cookie
			flushAuthCookie(c, cfg)

			redirectOrUnauthorized(c, unauthorizedPathRegexList)

			return
		}

		// Create new request with new context
		c.Request = c.Request.WithContext(
			SetAuthenticatedUserToContext(c.Request.Context(), &ouser),
		)
		// Add it to gin context
		SetAuthenticatedUserToGin(c, &ouser)

		logger.Infof("OIDC User authenticated: %s", ouser.GetIdentifier())
		c.Next()
	}
}

func flushAuthCookie(c *gin.Context, cfg *config.Config) {
	http.SetCookie(c.Writer, &http.Cookie{
		Expires:  time.Unix(0, 0),
		Name:     cfg.OIDCAuthentication.CookieName,
		Value:    "",
		HttpOnly: true,
		Secure:   cfg.OIDCAuthentication.CookieSecure,
		Path:     "/",
	})
}

func redirectOrUnauthorized(c *gin.Context, unauthorizedPathRegexList []*regexp.Regexp) {
	// Find a potential match into all regexps
	_, match := lo.Find(unauthorizedPathRegexList, func(reg *regexp.Regexp) bool {
		return reg.MatchString(c.Request.URL.Path)
	})

	if match {
		// Unauthorized error
		err := cerrors.NewUnauthorizedError("unauthorized")
		utils.AnswerWithError(c, err)

		return
	}

	// Initialize redirect URI
	rdURI := loginPath
	// Check if redirect URI must be created
	// If request path isn't equal to login path, build redirect URI to keep incoming request
	if c.Request.RequestURI != loginPath {
		// Build incoming request
		incomingURI := utils.GetRequestURL(c.Request)
		// URL Encode it
		urlEncodedIncomingURI := url.QueryEscape(incomingURI)
		// Build redirect URI
		rdURI = fmt.Sprintf("%s?%s=%s", loginPath, redirectQueryKey, urlEncodedIncomingURI)
	}

	// Redirect
	c.Redirect(http.StatusTemporaryRedirect, rdURI)
	c.Abort()
}

func getJWTToken(logger log.Logger, r *http.Request, cookieName string) (string, error) {
	logger.Debug("Try to get Authorization header from request")
	// Get Authorization header
	authHd := r.Header.Get("Authorization")
	// Check if Authorization header is populated
	if authHd != "" {
		// Split header to get token => Format "Bearer TOKEN"
		sp := strings.Split(authHd, " ")
		if len(sp) != 2 || sp[0] != "Bearer" {
			return "", errors.New("authorization header doesn't follow bearer format")
		}
		// Get content
		content := sp[1]
		// Check if content exists
		if content != "" {
			return content, nil
		}
	}
	// Content is empty => Try to continue with cookie

	logger.Debug("Try get auth cookie from request")
	// Try to get auth cookie
	cookie, err := r.Cookie(cookieName)
	// Check if error exists
	if err != nil {
		logger.Debug("Can't load auth cookie")

		if !errors.Is(err, http.ErrNoCookie) {
			return "", errors.WithStack(err)
		}
	}
	// Check if cookie value exists
	if cookie != nil {
		return cookie.Value, nil
	}

	return "", nil
}

// IsValidRedirect checks whether the redirect URL is whitelisted.
func isValidRedirect(redirectURLStr, reqURLStr string) (bool, error) {
	// Check if it isn't forged with complete urls
	if !strings.HasPrefix(redirectURLStr, "http://") &&
		!strings.HasPrefix(redirectURLStr, "https://") {
		return false, nil
	}

	// Parse request URL
	currentURL, err := url.Parse(reqURLStr)
	// Check error
	if err != nil {
		return false, errors.WithStack(err)
	}
	// Parse redirect URL
	redURL, err := url.Parse(redirectURLStr)
	// Check error
	if err != nil {
		return false, errors.WithStack(err)
	}

	// Check if hosts aren't the same
	if redURL.Host != currentURL.Host {
		return false, nil
	}

	// Default
	return true, nil
}
