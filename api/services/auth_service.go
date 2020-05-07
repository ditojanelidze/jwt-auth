package services

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/api/models"
	"github.com/ditojanelidze/jwt-auth/database"
)

func Authenticate(username string, password string) (map[string]string, error) {
	var tokens map[string]string
	db, err := database.Connect()
	if err!= nil{
		fmt.Println(err.Error())
		return tokens, err
	}
	defer db.Close()
	user := models.User{}
	err = db.Where(&models.User{Username: username, Password: password}).Find(&user).Error
	if err != nil{
		return tokens, err
	}
	tokens, err = RetrieveTokens(user)
	if err != nil{
		return tokens, err
	}
	return tokens, nil
}