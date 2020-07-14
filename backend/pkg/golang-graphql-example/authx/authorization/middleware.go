package authorization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

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
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})

			return
		}

		// Check if user is authorized
		if !authorized {
			err := fmt.Errorf("forbidden user %s", ouser.GetIdentifier())

			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err})

			return
		}

		// User is authorized

		logger.Infof("OIDC user %s authorized", ouser.GetIdentifier())
		c.Next()
	}
}
