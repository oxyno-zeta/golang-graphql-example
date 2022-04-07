package dataloaders

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business"
)

type contextKey struct {
	name string
}

var reqCtxKey = &contextKey{name: "graphql-dataloaders"}

func Middleware(busiSvr *business.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create dataloaders
		dl := new(busiSvr)

		// Update request with new context
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), reqCtxKey, dl))

		// Next
		c.Next()
	}
}

func GetDataloadersFromContext(ctx context.Context) *Dataloaders {
	return ctx.Value(reqCtxKey).(*Dataloaders) // nolint: forcetypeassert // Ignored
}
