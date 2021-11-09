package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
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

func TestLoginRoute(t *testing.T) {
	// TODO: add defer to delete user after test.
	userEmail := gofakeit.Email()
	userPassword := gofakeit.Password(true, true, true, true, false, 10)
	signUpData := gin.H{
		"email":            userEmail,
		"password":         userPassword,
		"password_confirm": userPassword,
	}

	signUpBody := convertJsonToBody(signUpData)

	w := request("POST", "/api/auth/v1/signup", signUpBody)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Run("should be logged in", func(t *testing.T) {
		loginData := gin.H{
			"email":    userEmail,
			"password": userPassword,
		}

		loginBody := convertJsonToBody(loginData)

		w := request("POST", "/api/auth/v1/login", loginBody)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should be failed to log in with wrong email", func(t *testing.T) {
		loginData := gin.H{
			"email":    gofakeit.Email(),
			"password": userPassword,
		}

		loginBody := convertJsonToBody(loginData)

		w := request("POST", "/api/auth/v1/login", loginBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should be failed to log in with wrong password", func(t *testing.T) {
		randomPassword := gofakeit.Password(true, true, true, true, false, 12)
		loginData := gin.H{
			"email":    userEmail,
			"password": randomPassword,
		}

		loginBody := convertJsonToBody(loginData)

		w := request("POST", "/api/auth/v1/login", loginBody)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestSignUpRoute(t *testing.T) {
	data := gin.H{
		"email":            "go.gin.signup@abc.com",
		"password":         "Abcd1234!@",
		"password_confirm": "Abcd1234!@",
	}

	t.Run("should be success when request body is correctly filled", func(t *testing.T) {
		body := convertJsonToBody(data)

		w := request("POST", "/api/auth/v1/signup", body)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	disallowedPasswordMap := map[string]string{
		"lowercase":         "ASDF1234!",
		"special_character": "Asdf1234",
		"uppercase":         "asdf1234!",
		"number":            "Asdf!@#$",
		"less_than_8":       "Aa1!",
	}

	for errName, disallowedPassword := range disallowedPasswordMap {
		t.Run(fmt.Sprintf("should be failed with weak password like %s", disallowedPassword), func(t *testing.T) {
			tmp := data
			tmp["password"] = disallowedPassword
			tmp["password_confirm"] = disallowedPassword
			body := convertJsonToBody(tmp)

			w := request("POST", "/api/auth/v1/signup", body)

			assert.Equal(t, http.StatusBadRequest, w.Code)

			resp := lib.ParseJSON(w.Body.Bytes())
			assert.Equal(t, errName, resp["error"])
		})
	}

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
