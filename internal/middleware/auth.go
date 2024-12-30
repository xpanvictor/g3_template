package middleware

import (
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func AuthMiddleware(auth0Domain, audience string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtValidator, err := validator.New(
			issuer: fmt.Sprintf("https://%s/", auth0Domain),
			audience: []string{audience},
	)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Validate token
		token := c.GetHeader("Authorization")
		// ... token validation logic
	}
}
