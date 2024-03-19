package helpers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "mygramhacktiv8"
var tokenExpirationDuration = time.Hour * 8

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"time":  time.Now().Add(tokenExpirationDuration).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}

	return signToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errRes := errors.New("token is invalid")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer ")

	if !bearer {
		return nil, errRes
	}

	tokenString := strings.TrimPrefix(headerToken, "Bearer ")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["time"].(float64))
		if time.Unix(exp, 0).Before(time.Now()) {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}

	return nil, errRes
}
