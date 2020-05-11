package services

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/api/models"
	"github.com/ditojanelidze/jwt-auth/database"
	"strconv"
)

func RefreshToken(token string) (map[string]string, error) {
	var tokens map[string]string
	stored_token := fmt.Sprintf("refresh_%s", token)
	redis_client := redisClient()
	user_id, err := redis_client.Get(stored_token).Result()
	if err != nil {
		return tokens, err
	}
	db, err := database.Connect()
	if err!= nil{
		fmt.Println(err.Error())
		return tokens, err
	}
	defer db.Close()
	user := models.User{}
	id, _ := strconv.Atoi(user_id)
	db.First(&user, id)
	tokens, err = Authenticate(user.Username,user.Password)
	redis_client.Del(stored_token)
	return tokens, nil
}
