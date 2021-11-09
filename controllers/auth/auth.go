package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"unicode"
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

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	// TODO: use go routine.
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasNumber == false {
		return "number"
	} else if hasLower == false {
		return "lowercase"
	} else if hasUpper == false {
		return "uppercase"
	} else if hasSpecial == false {
		return "special_character"
	} else {
		return ""
	}
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

	vulnerability := checkPasswordStrength(body.Password)
	if vulnerability != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": vulnerability})
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
