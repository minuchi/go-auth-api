package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func convertJsonToBody(data map[string]interface{}) *bytes.Buffer {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	return bytes.NewBuffer(b)
}

func TestHealthCheckRoute(t *testing.T) {
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

func TestSignUpRoute(t *testing.T) {
	data := gin.H{
		"email":            "go.gin@abc.com",
		"password":         "abcd1234!@",
		"password_confirm": "abcd1234!@",
	}

	t.Run("should be success when request body is correctly filled", func(t *testing.T) {
		body := convertJsonToBody(data)

		w := request("POST", "/api/auth/v1/signup", body)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	for key := range data {
		t.Run(fmt.Sprintf("should be failed when %s is not included in request body", key), func(t *testing.T) {
			tmp := data
			delete(tmp, key)
			body := convertJsonToBody(tmp)

			w := request("POST", "/api/auth/v1/signup", body)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}
