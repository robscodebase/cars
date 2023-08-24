package middleware

import (
	"cars/store/redis"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	tokenKey               = "token"
	authorizationHeaderKey = "Authorization"
)

func Authenticate(redisClient *redis.Redis) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		authHeader := gctx.Request.Header.Get(authorizationHeaderKey)
		if authHeader == "" {
			gctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			gctx.Abort()
			return
		}

		// Check for Bearer schema
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			gctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is Bearer (token)"})
			gctx.Abort()
			return
		}

		tokenString := parts[1]

		exists, err := redisClient.TokenExists(tokenString)
		if err != nil || exists == 0 {
			gctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found or invalid"})
			gctx.Abort()
			return
		}

		setTokenInContext(gctx, tokenString)
		gctx.Next()
	}
}

func setTokenInContext(gctx *gin.Context, token string) {
	gctx.Set(tokenKey, token)
}

func GetTokenFromContext(gctx *gin.Context) string {
	token, _ := gctx.Get(tokenKey)
	return token.(string)
}
