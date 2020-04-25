package initializers

import (
	"fmt"
	"os"
)

var (
	DB_DRIVER = "mysql"
	DBURL     = ""
)

func DbInit() {
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))
}
