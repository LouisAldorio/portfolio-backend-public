package config

import (
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *gorm.DB{
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open("mysql", connection_string)

	if err != nil {
		fmt.Print(err)
		panic("Failed to connect database")
	}

	db.LogMode(true)
	db.SingularTable(true)

	return db
}