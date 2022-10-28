package database

import (
	"be_project3team3/config"
	rOrder "be_project3team3/feature/order/repository"
	rProduct "be_project3team3/feature/product/repository"
	rUser "be_project3team3/feature/user/repository"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *config.AppConfig) *gorm.DB {
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
	db.AutoMigrate(&rUser.User{})
	db.AutoMigrate(&rProduct.Product{})
	db.AutoMigrate(&rOrder.Order{})
}
