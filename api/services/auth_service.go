package services

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/api/models"
	"github.com/ditojanelidze/jwt-auth/database"
)

func Authenticate(username string, password string) (string, error) {
	var token string
	db, err := database.Connect()
	if err!= nil{
		fmt.Println(err.Error())
		return token, err
	}
	defer db.Close()
	user := models.User{}
	err = db.Where(&models.User{Username: username, Password: password}).Take(&models.User{}).Error
	if err != nil{
		return token, err
	}
	token, err = RetrieveTokens(user)
	if err != nil{
		return token, err
	}
	return token, nil
}