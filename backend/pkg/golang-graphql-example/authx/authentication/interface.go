package authentication

import (
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
)

//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication Service
type Service interface {
	// Middleware will redirect authentication to basic auth or OIDC depending on request path and resources declared.
	Middleware(unauthorizedPathRegexList []*regexp.Regexp) gin.HandlerFunc
	// OIDCEndpoints will set OpenID Connect endpoints for authentication and callback.
	OIDCEndpoints(router gin.IRouter) error
}

func NewService(cfgManager config.Manager) Service {
	return &service{
		cfgManager: cfgManager,
	}
}
