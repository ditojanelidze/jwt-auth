package initializers

import "os"

var (
	REDIS_URL  = "mysql"
	REDIS_PORT = ""
)

func RedisInit() {
	REDIS_URL = os.Getenv("REDIS_URL")
	REDIS_PORT = os.Getenv("REDIS_PORT")
}
