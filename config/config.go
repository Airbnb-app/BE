package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     uint
	DB_NAME     string
	SERVER_PORT uint
	JWT_SECRET  string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	if _, exist := os.LookupEnv("SECRET"); !exist {
		if err := godotenv.Load(".env"); err != nil {
			log.Println(err)
		}
	}

	// SECRET = os.Getenv("SECRET")
	cnvServerPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Cannot parse Server Port variable")
		return nil
	}
	defaultConfig.SERVER_PORT = uint(cnvServerPort)
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	cnvDBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.DB_PORT = uint(cnvDBPort)
	defaultConfig.JWT_SECRET = os.Getenv("JWT_SECRET")

	return &defaultConfig
}
