package middleware

import (
	redis2 "cars/store/redis"
	tests2 "cars/util/tests"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const testTokenSecret = "7cP0iH2EfUv9y$B^5p4s%#jL0kZvV@r&"

func TestAuthenticate(t *testing.T) {
	tests2.SetTestEnvVars()
	redis, err := redis2.NewRedis()
	require.NoError(t, err)
	require.NotNil(t, redis)
	validToken := generateMockJWTToken()
	err = redis.SetToken(validToken)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "missing Authorization header",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Authorization header not provided",
		},
		{
			name:           "incorrect Authorization format",
			authHeader:     "BearerTokenWithoutSpace",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Authorization header format is Bearer (token)",
		},
		{
			name:           "invalid token",
			authHeader:     "Bearer invalidToken",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Token not found or invalid",
		},
		{
			name:           "valid token",
			authHeader:     "Bearer " + validToken,
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.Use(Authenticate(redis))
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "OK")
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set(authorizationHeaderKey, tt.authHeader)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus == http.StatusUnauthorized {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

func generateMockJWTToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte(testTokenSecret))
	return tokenString
}
