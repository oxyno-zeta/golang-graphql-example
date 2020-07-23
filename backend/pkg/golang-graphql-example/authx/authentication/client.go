package authentication

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
)

type Client interface {
	// Middleware will redirect authentication to basic auth or OIDC depending on request path and resources declared
	Middleware(unauthorizedPathRegexList []*regexp.Regexp) gin.HandlerFunc
	// OIDCEndpoints will set OpenID Connect endpoints for authentication and callback
	OIDCEndpoints(router gin.IRouter) error
}

func NewService(cfgManager config.Manager) Client {
	return &service{
		cfgManager: cfgManager,
	}
}
