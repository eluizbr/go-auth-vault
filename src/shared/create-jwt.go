package shared

import (
	"time"

	"github.com/eluizbr/go_pagamentos/auth/src/auth/models"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(user models.User) (map[string]string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"name": user.Name,
		"id":   user.ID,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": t,
	}, nil

}
