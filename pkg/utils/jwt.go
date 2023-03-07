package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("my_secret_key")

func GenerateJWT(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.GetHeader("Authorization")
		if userToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "로그인이 필요한 서비스입니다."})
			return
		} // "Bearer <token>" 형식으로 토큰이 전달되기 때문에 " "로 분리하여 실제 토큰 값만 추출
		splitToken := strings.Split(userToken, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "올바르지 않은 토큰입니다."})
			return
		}

		tokenString := splitToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "올바르지 않은 토큰입니다."})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			currentUserId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userId"]), 10, 64)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "올바르지 않은 토큰입니다."})
				return
			}
			c.Set("currentUserId", currentUserId)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "올바르지 않은 토큰입니다."})
		return
	}
}
