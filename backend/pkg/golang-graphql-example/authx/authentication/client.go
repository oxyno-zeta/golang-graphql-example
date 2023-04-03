package authentication

import (
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
)

//go:generate mockgen -destination=./mocks/mock_Client.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication Client
type Client interface {
	// Middleware will redirect authentication to basic auth or OIDC depending on request path and resources declared.
	Middleware(unauthorizedPathRegexList []*regexp.Regexp) gin.HandlerFunc
	// OIDCEndpoints will set OpenID Connect endpoints for authentication and callback.
	OIDCEndpoints(router gin.IRouter) error
}

type providerEndpointsClaims struct {
	EndSessionEndpointURL *url.URL
	EndSessionEndpoint    string `json:"end_session_endpoint"`
}

func NewService(cfgManager config.Manager) Client {
	return &service{
		cfgManager: cfgManager,
	}
}
