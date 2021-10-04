package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type getTimeResponse struct {
	Time string `json:"time"`
}

func returnOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func GetTime(c *gin.Context) {
	t := &getTimeResponse{
		Time: time.Now().Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, t)
}

func Login(c *gin.Context) {
	returnOK(c)
}

func SignUp(c *gin.Context) {
	returnOK(c)
}

func IssueAccessToken(c *gin.Context) {
	returnOK(c)
}

func Verify(c *gin.Context) {
	returnOK(c)
}
