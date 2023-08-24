package tests

import (
	"bytes"
	"cars/store/redis"
	"cars/util"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetAuthRequest(t *testing.T, server *httptest.Server, redis *redis.Redis, method, endpoint string, body interface{}) *http.Request {
	token, err := util.GenerateJWTToken()
	require.NotEmptyf(t, token, "token is empty: %v", err)
	require.NoError(t, err)
	err = redis.SetToken(token)
	payload, err := json.Marshal(body)
	require.NoError(t, err)
	req, err := http.NewRequest(method, server.URL+endpoint, bytes.NewBuffer(payload))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	return req
}
