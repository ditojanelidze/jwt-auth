package database

import (
	"fmt"
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() (*gorm.DB, error){
	fmt.Println(initializers.DBURL)
	return gorm.Open(initializers.DB_DRIVER, initializers.DBURL)
}

