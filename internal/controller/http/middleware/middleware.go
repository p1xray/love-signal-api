package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	jwtmiddleware "github.com/p1xray/pxr-sso/pkg/jwt"
	"github.com/p1xray/pxr-sso/pkg/jwt/validator"
	"log"
	"love-signal-api/internal/server"
	"net/http"
	"os"
)

var (
	keyFunc = func(ctx context.Context) ([]byte, error) {
		clientSecret := os.Getenv("CLIENT_SECRET")
		return []byte(clientSecret), nil
	}
)

func CheckJWT() gin.HandlerFunc {
	issuer := os.Getenv("ISSUER")
	audience := os.Getenv("AUDIENCE")

	jwtValidator, err := validator.New(
		keyFunc,
		issuer,
		[]string{audience},
	)
	if err != nil {
		log.Fatalf("error creating JWT validator: %v", err)
	}

	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	return func(c *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			c.Request = r
			c.Next()
		}

		middleware.ParseJWT(handler).ServeHTTP(c.Writer, c.Request)

		if encounteredError {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "JWT is invalid."},
			)
		}
	}
}

func HasScope(expectedScope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userHasScope, err := server.UserHasScope(c, expectedScope)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				map[string]string{"message": err.Error()},
			)
		}

		if !userHasScope {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				map[string]string{"message": "User does not have the required permission."},
			)
		}

		c.Next()
	}
}
