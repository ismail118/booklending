package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ismail118/booklending/token"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "authorization"
	authorizationType   = "bearer"
)

var ErrNoAccessToken = errors.New("no access token found")
var ErrUnsupportedAuthType = errors.New("unsupported auth type")

func authMiddleware(paseto *token.Paseto) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(authorizationHeader)
		if len(accessToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrNoAccessToken))
			return
		}

		fields := strings.Fields(accessToken)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrNoAccessToken))
			return
		}

		if authorizationType != strings.ToLower(fields[0]) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrUnsupportedAuthType))
			return
		}

		tkn := fields[1]
		_, err := paseto.VerifyToken(tkn)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		c.Next()
	}
}
