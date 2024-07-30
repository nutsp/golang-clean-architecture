package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/nutsp/golang-clean-architecture/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
