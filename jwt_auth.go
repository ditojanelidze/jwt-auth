package main

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/config"
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/ditojanelidze/jwt-auth/server"
)

func main() {
	config.Init()
	server.Run()
	fmt.Println(initializers.DB_DRIVER)
	fmt.Println(initializers.DBURL)
	fmt.Println(initializers.REDIS_PORT)
	fmt.Println(initializers.REDIS_URL)
}
