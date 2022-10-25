package database

import (
	"ecommerce/config"
	"fmt"

	// "os"

	checkout "ecommerce/features/checkout/repository"
	user "ecommerce/features/user/repository"

	// "fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	c := config.NewConfig()
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPwd,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error :", err.Error())
		return nil
	}
	migrateDB(db)
	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.Product{})
	db.AutoMigrate(&user.Cart{})
	db.AutoMigrate(&checkout.Checkout{})
	db.AutoMigrate(&checkout.History{})
}
