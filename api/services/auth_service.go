package services

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/api/models"
	"github.com/ditojanelidze/jwt-auth/database"
)

func Authenticate(username string, password string) (map[string]string, error) {
	fmt.Println("111111111111")
	var tokens map[string]string
	db, err := database.Connect()
	if err!= nil{
		fmt.Println("22222222222222")
		fmt.Println(err.Error())
		return tokens, err
	}
	fmt.Println("33333333333")
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