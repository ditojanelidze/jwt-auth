package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ditojanelidze/jwt-auth/api/models"
	"os"
	"time"
)

func RetrieveTokens(user models.User) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp_date": time.Now().Add(time.Hour * 1),
	})
	tokenString, err := token.SignedString(os.Getenv("API_SECRET_KEY"))
	return tokenString, err
}
