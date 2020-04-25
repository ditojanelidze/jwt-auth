package server

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/config"
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/ditojanelidze/jwt-auth/router"
	"net/http"
)

func Run(){
	config.Init()
	listen()
}

func listen(){
	r := router.Router()
	fmt.Printf("\n\tListening [::]:%d\n\n", initializers.PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", initializers.PORT), r)
}