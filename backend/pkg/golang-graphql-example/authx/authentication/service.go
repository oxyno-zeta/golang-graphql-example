package authentication

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/thoas/go-funk"
	"golang.org/x/oauth2"
)

const callbackPath = "/auth/oidc/callback"
const loginPath = "/auth/oidc"
const userContextKeyName = "USER_CONTEXT_KEY"
const redirectQueryKey = "rd"

var userContextKey = &contextKey{name: userContextKeyName}

type service struct {
	verifier   *oidc.IDTokenVerifier
	cfgManager config.Manager
}

// GetAuthenticatedUser will get authenticated user in context
func GetAuthenticatedUserFromContext(ctx context.Context) *models.OIDCUser {
	res, _ := ctx.Value(userContextKey).(*models.OIDCUser)
	return res
}

// GetAuthenticatedUser will get authenticated user in context
func GetAuthenticatedUserFromGin(c *gin.Context) *models.OIDCUser {
	res, _ := c.Get(userContextKeyName)
	res1 := res.(*models.OIDCUser)

	return res1
}

func buildOauthRedirectURIParam(mainRedirectURLStr, rdVal string) (oauth2.AuthCodeOption, error) {
	oidcRedirectURL, err := url.Parse(mainRedirectURLStr)
	// Check if error exists
	if err != nil {
		return nil, err
	}
	// Get query params
	qsValues := oidcRedirectURL.Query()
	// Check if redirect value exists
	if rdVal != "" {
		// Add query param
		qsValues.Add(redirectQueryKey, rdVal)
	}
	// Add query params to oidc redirect url
	oidcRedirectURL.RawQuery = qsValues.Encode()
	// Build new oidc redirect url string
	oidcRedirectURLStr := oidcRedirectURL.String()

	authP := oauth2.SetAuthURLParam("redirect_uri", oidcRedirectURLStr)

	return authP, nil
}

// OIDCEndpoints will set OpenID Connect endpoints for authentication and callback
func (s *service) OIDCEndpoints(router gin.IRouter) error {
	ctx := context.Background()

	// Get configuration
	cfg := s.cfgManager.GetConfig()

	provider, err := oidc.NewProvider(ctx, cfg.OIDCAuthentication.IssuerURL)
	if err != nil {
		return err
	}

	oidcConfig := &oidc.Config{
		ClientID: cfg.OIDCAuthentication.ClientID,
	}
	verifier := provider.Verifier(oidcConfig)

	// Build redirect url
	mainRedirectURLObject, err := url.Parse(cfg.OIDCAuthentication.RedirectURL)
	// Check if error exists
	if err != nil {
		return err
	}
	// Continue to build redirect url
	mainRedirectURLObject.Path = path.Join(mainRedirectURLObject.Path, callbackPath)
	mainRedirectURLStr := mainRedirectURLObject.String()

	// Create OIDC configuration
	config := oauth2.Config{
		ClientID:    cfg.OIDCAuthentication.ClientID,
		Endpoint:    provider.Endpoint(),
		Scopes:      cfg.OIDCAuthentication.Scopes,
		RedirectURL: mainRedirectURLStr,
	}
	if cfg.OIDCAuthentication.ClientSecret != nil {
		config.ClientSecret = cfg.OIDCAuthentication.ClientSecret.Value
	}

	// Store state
	state := cfg.OIDCAuthentication.State

	// Store provider verifier in map
	s.verifier = verifier

	router.GET(loginPath, func(c *gin.Context) {
		// Get logger from request
		logger := log.GetLoggerFromGin(c)
		// Get redirect query from query params
		rdVal := c.Query(redirectQueryKey)
		// Need to build new oidc redirect url
		authParam, err := buildOauthRedirectURIParam(mainRedirectURLStr, rdVal)
		// Check if error exists
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.Redirect(http.StatusFound, config.AuthCodeURL(state, authParam))
		c.Abort()
	})

	router.GET(mainRedirectURLObject.Path, func(c *gin.Context) {
		// Get logger from request
		logger := log.GetLoggerFromGin(c)

		// Get redirect url
		rdVal := c.Query(redirectQueryKey)
		// Check if rdVal exists and that redirect url value is valid
		if rdVal != "" && !isValidRedirect(rdVal) {
			err := errors.New("redirect url is invalid")

			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// Check state
		if c.Query("state") != state {
			err := errors.New("state did not match")
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// Build auth param
		authParam, err := buildOauthRedirectURIParam(mainRedirectURLStr, rdVal)
		// Check if error exists
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		oauth2Token, err := config.Exchange(ctx, c.Query("code"), authParam)
		if err != nil {
			err = errors.New("failed to exchange token: " + err.Error())
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			err = errors.New("no id_token field in token")
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			err = errors.New("failed to verify ID Token: " + err.Error())
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		var resp models.OIDCUser

		// Try to open JWT token in order to verify that we can open it
		err = idToken.Claims(&resp)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		// Now, we know that we can open jwt token to get claims

		// Build cookie
		cookie := &http.Cookie{
			Expires:  oauth2Token.Expiry,
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})

			return
		}
		// Check if JWT content is empty or not
		if jwtContent == "" {
			logger.Error("No auth header or cookie detected, redirect to oidc login")
			s.redirectOrAuthorized(c, unauthorizedPathRegexList)

			return
		}

		// Parse token

		var ouser models.OIDCUser
		// Verify token
		idToken, err := s.verifier.Verify(context.Background(), jwtContent)
		// Check error
		if err != nil {
			logger.Error(err)
			// Flush potential cookie
			http.SetCookie(c.Writer, &http.Cookie{
				Expires:  time.Unix(0, 0),
				Name:     cfg.OIDCAuthentication.CookieName,
				Value:    "",
				HttpOnly: true,
				Secure:   cfg.OIDCAuthentication.CookieSecure,
				Path:     "/",
			})

			s.redirectOrAuthorized(c, unauthorizedPathRegexList)

			return
		}

		// Get claims
		err = idToken.Claims(&ouser)
		if err != nil {
			logger.Error(err)
			// Flush potential cookie
			http.SetCookie(c.Writer, &http.Cookie{
				Expires:  time.Unix(0, 0),
				Name:     cfg.OIDCAuthentication.CookieName,
				Value:    "",
				HttpOnly: true,
				Secure:   cfg.OIDCAuthentication.CookieSecure,
				Path:     "/",
			})

			s.redirectOrAuthorized(c, unauthorizedPathRegexList)

			return
		}

		// Add user to request context by creating a new context
		ctx := context.WithValue(c.Request.Context(), userContextKey, &ouser)
		// Create new request with new context
		c.Request = c.Request.WithContext(ctx)
		// Add it to gin context
		c.Set(userContextKeyName, &ouser)

		logger.Infof("OIDC User authenticated: %s", ouser.GetIdentifier())
		c.Next()
	}
}

func (s *service) redirectOrAuthorized(c *gin.Context, unauthorizedPathRegexList []*regexp.Regexp) {
	// Find a potential match into all regexps
	match := funk.Find(unauthorizedPathRegexList, func(reg *regexp.Regexp) bool {
		return reg.MatchString(c.Request.URL.Path)
	})

	if match != nil {
		// Unauthorized error
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
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

		if err != http.ErrNoCookie {
			return "", err
		}
	}
	// Check if cookie value exists
	if cookie != nil {
		return cookie.Value, nil
	}

	return "", nil
}

// IsValidRedirect checks whether the redirect URL is whitelisted
func isValidRedirect(redirect string) bool {
	return strings.HasPrefix(redirect, "http://") || strings.HasPrefix(redirect, "https://")
}
