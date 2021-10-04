package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minuchi/go-auth-api/lib"
	"github.com/stretchr/testify/assert"
)

func request(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	router := lib.SetupRouter(gin.TestMode)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	router.ServeHTTP(w, req)

	return w
}

func TestHealthzRoute(t *testing.T) {
	w := request("GET", "/healthz", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	resp := lib.ParseJSON(w.Body.Bytes())
	assert.Equal(t, true, resp["ok"])
}

func TestGetTimeRoute(t *testing.T) {
	now := time.Now()
	w := request("GET", "/api/auth/v1/time", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	resp := lib.ParseJSON(w.Body.Bytes())

	respTime := fmt.Sprintf("%s", resp["time"])

	resultTime, err := time.Parse(time.RFC3339, respTime)
	assert.NoError(t, err)

	diff := now.Sub(resultTime)
	assert.Less(t, diff, time.Second*2)
}
