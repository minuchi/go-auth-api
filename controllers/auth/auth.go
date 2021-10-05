package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type signupRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

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
	var body signupRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: insert a row to mail table
	if mode, _ := c.Get("mode"); mode != gin.TestMode {
	}

	returnOK(c)
}

func IssueAccessToken(c *gin.Context) {
	returnOK(c)
}

func Verify(c *gin.Context) {
	returnOK(c)
}
