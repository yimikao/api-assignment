package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/yimikao/api-assignment/api/models"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dialect := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")

	return &Config{
		DB: &DBConfig{
			Dialect:  dialect,
			Username: username,
			Password: password,
			Name:     name,
			Charset:  charset,
		},
	}
}

var DB *gorm.DB

//Initialize database with the necessary configurations
func InitDB(c *Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		c.DB.Username,
		c.DB.Password,
		c.DB.Name,
		c.DB.Charset)

	db, err := gorm.Open("mysql", dbURI)

	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
	DB.AutoMigrate(&models.Profile{})
	log.Println("Database migrated successfully")

}
