package initializers

import (
	"os"
	"strconv"
)

var (
	REDIS_URL  = ""
	REDIS_PORT = ""
	REDIS_DB = 0
)

func RedisInit() {
	REDIS_URL = os.Getenv("REDIS_URL")
	REDIS_PORT = os.Getenv("REDIS_PORT")
	REDIS_DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
}
