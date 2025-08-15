package authorization

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	goerrors "emperror.dev/errors"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type restInputOPA struct {
	Input *restInputDataOPA `json:"input"`
}

type restInputDataOPA struct {
	User    *models.OIDCUser    `json:"user"`
	Request *httpRequestDataOPA `json:"request"`
	Tags    map[string]string   `json:"tags"`
}

type httpRequestDataOPA struct {
	Headers    map[string]string `json:"headers"`
	Method     string            `json:"method"`
	Protocol   string            `json:"protocol"`
	RemoteAddr string            `json:"remoteAddr"`
	Scheme     string            `json:"scheme"`
	Host       string            `json:"host"`
	Path       string            `json:"path"`
	ParsedPath []string          `json:"parsed_path"`
}

func (s *service) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get logger
		logger := log.GetLoggerFromGin(c)
		// Get user from request
		ouser := authentication.GetAuthenticatedUserFromGin(c)

		authorized, err := s.isRequestAuthorized(c.Request, ouser)
		// Check error
		if err != nil {
			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		// Check if user is authorized
		if !authorized {
			err = errors.NewForbiddenError("forbidden")

			logger.Error(err)
			utils.AnswerWithError(c, err)

			return
		}

		// User is authorized

		logger.Infof("OIDC user %s authorized", ouser.GetIdentifier())
		c.Next()
	}
}

func (s *service) isRequestAuthorized(req *http.Request, oidcUser *models.OIDCUser) (bool, error) {
	// Get configuration
	opaServerCfg := s.cfgManager.GetConfig().OPAServerAuthorization

	// Transform headers into map
	headers := make(map[string]string)
	for k, v := range req.Header {
		headers[strings.ToLower(k)] = v[0]
	}
	// Parse path
	parsedPath := deleteEmpty(strings.Split(req.RequestURI, "/"))
	// Calculate scheme
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	// Generate OPA Server input data
	input := &restInputOPA{
		Input: &restInputDataOPA{
			User: oidcUser,
			Tags: opaServerCfg.Tags,
			Request: &httpRequestDataOPA{
				Method:     req.Method,
				Protocol:   req.Proto,
				Headers:    headers,
				RemoteAddr: req.RemoteAddr,
				Scheme:     scheme,
				Host:       req.Host,
				ParsedPath: parsedPath,
				Path:       req.RequestURI,
			},
		},
	}
	// Json encode body
	bb, err := json.Marshal(input)
	if err != nil {
		return false, goerrors.WithStack(err)
	}

	return s.requestOPAServer(req.Context(), opaServerCfg, bb)
}

func deleteEmpty(s []string) []string {
	var r []string

	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}
