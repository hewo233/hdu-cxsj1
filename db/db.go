package db

import (
	"fmt"
	"github.com/hewo233/hdu-cxsj1/module"
	"github.com/hewo233/hdu-cxsj1/shared/consts"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func CreateDB() {
	err := DB.AutoMigrate(&module.User{})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func ConnectDB() {

	if err := godotenv.Load(consts.EnvFile); err != nil {
		log.Println("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
}
