package config

import (
	"fmt"
	"os"

	m "github.com/forumGamers/payment-service/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func Connection(){
	if err := godotenv.Load() ; err != nil {
		panic(err.Error())
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")

	if DATABASE_URL == "" {
		DATABASE_URL = "user=postgres password=qwertyui host=127.0.0.1 port=5432 dbname=payment-service sslmode=disable"
	}

	database,err := gorm.Open("postgres",DATABASE_URL)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("connection success")

	database.AutoMigrate(&m.Transaction{})
}