package middleware

import (
	"errors"
	"fmt"
	"gin-api-test/api/v1/models"
	"gin-api-test/helpers/jwt_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CekUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		MiddlewareImp(c)
	}
}

func MiddlewareImp(c *gin.Context) {
	token, err := GetTokenFromHeader(c)
	if err != nil {
		c.Header("WWW-Authenticate", "JWT realm=jwt-token")
		c.Abort()
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "failed", "UnAutorization", nil))
		return
	}

	getUserId, err := jwt_token.ValidateTokenHeader(token)
	if err != nil {
		c.Header("WWW-Authenticate", "JWT realm=jwt-token")
		c.Abort()
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "failed", "UnAutorization", nil))
		return
	}

	c.Set("userId", getUserId)
	c.Set("token", token)
}

func GetTokenFromHeader(c *gin.Context) (string, error) {
	fmt.Println("ini awal get token")
	authHeader := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.Set("ResCode", "400")
		return "", errors.New("Unknown Auth Bearer")
	}
	fmt.Println("parser token ", parts[1])
	return parts[1], nil
}
