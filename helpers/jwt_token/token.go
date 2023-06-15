package jwt_token

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gin-api-test/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodePrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("Error Decode PrivateKey :%w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodePrivateKey)
	if err != nil {
		return "", fmt.Errorf("Error create parse key:%w", err)
	}

	fmt.Println("parse private", key)
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["user_id"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create token error: %w", err)
	}

	return token, nil

}

func ValidateToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("coult not decode:%w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate parse key:%w", err)
	}

	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpacted Method: %s", t.Header["exp"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate :%w", err)
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok || !parseToken.Valid {
		return nil, fmt.Errorf("validate:invalid token")
	}

	return claims["user_id"], nil
}

func ValidateTokenHeader(token string) (string, error) {
	decoderPublicKey, err := base64.StdEncoding.DecodeString(config.MyEnv.AccTokenPublicKey)
	if err != nil {
		return " ", fmt.Errorf("Could not decode: %w", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(decoderPublicKey)
	if err != nil {
		return " ", fmt.Errorf("validate parse key: %w", err)
	}
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			fmt.Println("ini oke", ok)

			return nil, fmt.Errorf("Unexpacted Method: %s", t.Header["exp"])
		}
		return key, nil
	})
	if err != nil {

		return " ", fmt.Errorf("validate :%w", err)
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok || !parseToken.Valid {
		return " ", fmt.Errorf("validate:invalid token")
	}
	return claims["user_id"].(string), nil
}

func GetUserID(c *gin.Context) (string, error) {

	uid, valid := c.Get("userId")
	if !valid {
		return "", errors.New("undifined User ID")
	}
	return uid.(string), nil
}
