package db

import (
	"fmt"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sumitalp/productcatalog/models"
)

func New() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@/adcash?charset=utf8&parseTime=True&loc=Local", 
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD")
		))
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Tag{},
	)
}
