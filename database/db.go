package database

import (
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error){
	return gorm.Open(initializers.DB_DRIVER, initializers.DBURL)
}

