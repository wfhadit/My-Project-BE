package config

import (
	"fmt"
	"my-project-be/features/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var JWTSECRET = ""

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

func assignEnv(c AppConfig) (AppConfig, bool) {
	missing := false
	if val, found := os.LookupEnv("DBUsername"); found {
		c.DBUsername = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPassword"); found {
		c.DBPassword = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPort"); found {
		c.DBPort = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBHost"); found {
		c.DBHost = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBName"); found {
		c.DBName = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("JWTSECRET"); found {
		JWTSECRET = val
	} else {
		missing = true
	}
	return c, missing
}

func InitConfig() AppConfig {
	result := AppConfig{}
	missing := false
	result, missing = assignEnv(result)
	if missing {
		godotenv.Load(".env")
		result, _ = assignEnv(result)
	}

	return result
}

func InitSQL(c AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("terjadi error", err.Error())
		return nil
	}

	db.AutoMigrate(&user.User{})

	return db
}