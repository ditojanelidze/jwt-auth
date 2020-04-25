package initializers

import (
	"log"
	"os"
	"strconv"
)

var (
	PORT  = 9000
)

func InitServer() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println("Warning! \nUnable to determine API_PORT enviroemt parameter, running on default port - 9000")
	}
}
