package crypto

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

type DataToken struct {
	Id   int
	Name string
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(jwt.StandardClaims{
		Id:        "123",
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(), // Set token expiration time
		IssuedAt:  time.Now().Unix(),                      // Set token issuance time
		Subject:   "TOKEN",                                // TOKEN | REFRESH_TOKEN
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("token : ", token)
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk0ODMzOTAsImp0aSI6IjMiLCJpYXQiOjE2ODk0NDczOTAsInN1YiI6IlRPS0VOIn0.h76Et8jeNdZetvtWOrX90I5LK98eGBNp4OI9KaLgLVQ"
	claims, err := ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("claims : ", claims.Id)
}
