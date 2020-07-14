package authorization

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
)

type Service interface {
	Middleware() gin.HandlerFunc
	IsAuthorized(ctx context.Context, action, resource string) (bool, error)
}

func NewService(cfgManager config.Manager) Service {
	return &service{
		cfgManager: cfgManager,
	}
}
