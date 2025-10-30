package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dborowsky/simplebank/token"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "authorization"
	ContentTypeBearer       = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header is malformed")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != ContentTypeBearer {
			err := fmt.Errorf("authorization type is not supported: %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
