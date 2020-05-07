package services

import "fmt"

func LogOut(token string) {
	redis := redisClient()
	redis.Del(fmt.Sprintf("refresh_%s",token))
}