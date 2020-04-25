package services

import (
	"github.com/ditojanelidze/jwt-auth/api/models"
	"github.com/ditojanelidze/jwt-auth/database"
)

func AuthService(username string, password string) (string, error) {
	var token string
	db, err := database.Connect()
	if err!= nil{
		return "", err
	}
	defer db.Close()
	user := models.User{}
	err = db.Where(&models.User{Username: username, Password: password}).Take(&models.User{}).Error
	if err != nil{
		return "", err
	}
	token, err = RetrieveTokens(user)
	if err != nil{
		return "", err
	}
	return token, nil
}
