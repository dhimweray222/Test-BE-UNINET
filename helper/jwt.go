package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
)

type Claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJwt(id, email string) (string, error) {
	//jwt
	SecretKey := os.Getenv("SECRET_KEY")

	SessionLogin := os.Getenv("SESSION_LOGIN")

	session, _ := strconv.Atoi(SessionLogin)
	claims := &Claims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(session)).Unix(),
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokens.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, string, error) {
	//jwt
	SecretKey := os.Getenv("SECRET_KEY")
	var claims Claims
	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}

	return claims.Id, claims.Email, nil
}
