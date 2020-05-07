package services

import "fmt"

func LogOut(token string) {
	fmt.Println(token)
	redis := redisClient()
	redis.Del(fmt.Sprintf("refresh_%s",token))
}