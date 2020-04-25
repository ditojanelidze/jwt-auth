package database

import (
	"github.com/ditojanelidze/jwt-auth/config/initializers"
	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error){
	return gorm.Open(initializers.DB_DRIVER, initializers.DBURL)
}

