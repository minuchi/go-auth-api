package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func returnOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func GetTime(c *gin.Context) {
	returnOK(c)
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
