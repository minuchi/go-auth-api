package auth

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
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

func checkPasswordStrength(password string) string {
	const MinimumPasswordLength = 8
	length := len(password)
	if length < MinimumPasswordLength {
		return fmt.Sprintf("less_than_%d", MinimumPasswordLength)
	}

	// TODO: All special characters should be included.
	patterns := map[string]string{
		"01|number":            "[0-9]",
		"02|lowercase":         "[a-z]",
		"03|uppercase":         "[A-Z]",
		"04|special_character": "[!@#$%^&:;?]",
	}

	for name, pattern := range patterns {
		result, _ := regexp.MatchString(pattern, password)
		if result == false {
			return strings.Split(name, "|")[1]
		}
	}

	return ""
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

	weakName := checkPasswordStrength(body.Password)
	if weakName != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": weakName})
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
