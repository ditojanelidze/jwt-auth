package services

import (
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/dgrijalva/jwt-go"
	"github.com/ditojanelidze/jwt-auth/api/models"
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/go-redis/redis/v7"
	"os"
	"time"
)
var (
	EXPIRATION = time.Hour * 2
	REFRESH_LENGTH = 20
)
func RetrieveTokens(user models.User) (map[string]string, error){
	tokenString, err := generateAccessToken(user.ID)
	if err != nil{
		return map[string]string{}, err
	}
	refreshToken := uniuri.NewLen(REFRESH_LENGTH)
	tokens := map[string]string {
		"access_token": tokenString,
		"refresh_token": refreshToken,
	}
	storeInRedis(user, tokens)
	return tokens, err
}

func generateAccessToken(id int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp_date": time.Now().Add(time.Minute * 20),
	})
	return token.SignedString([]byte(os.Getenv("API_SECRET_KEY")))
}

func storeInRedis(user models.User, tokens map[string]string) {
	fmt.Println(user)
	redisClient := redis.NewClient(&redis.Options{
		//fmt.Sprintf("%s:%s",initializers.REDIS_URL, initializers.REDIS_PORT)
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: initializers.REDIS_DB,
	})
	key := fmt.Sprintf("%s_%d", "auth_tokens", user.ID)
	redisClient.Set(key, tokens, EXPIRATION)
}