package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ditojanelidze/jwt-auth/api/models"
	"time"
)

func RetrieveTokens(user models.User) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp_date": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	return token
}
