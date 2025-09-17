package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("THIS_IS_SECRET_KEY")
	claims := jwt.MapClaims{
		"uid": "800593256",
		"role": "master",
		"exp": time.Now().Add(time.Hour).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedtoken, err := token.SignedString(secret)

	if err != nil {
		panic(err)
	}

	fmt.Println("Test Generate JWT Token")
	fmt.Println("Hasil token: ", signedtoken)
}