package initializers

import (
	"fmt"
	"os"
)

var (
	DB_DRIVER = ""
	DBURL     = ""
)

func DbInit() {
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),)
}
