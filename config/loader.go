package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	C    *Config // Singleton
	once sync.Once
)

func Load(path string) {

	once.Do(func() {

		if path == "" {
			log.Fatalln("path is empty")
		}

		if path == "http_request" || path == "worker" {
			getwd, err := os.Getwd()
			if err != nil {
				log.Fatalln("error getting current working directory :", err)
			}

			path = filepath.Join(getwd, ".env")
		} else if path == "config" {
			path = filepath.Join("..", ".env")

		}

		err := godotenv.Load(path)
		if err != nil {
			log.Fatalln("Error loading .env file  , error:", err)
		}

		// set global config
		C = &Config{
			App: AppConfig{
				Name:  getString("APP_NAME", "myapp"),
				Env:   getString("APP_ENV", "dev"),
				Debug: getBool("DEBUG", true),
				Port:  getInt("APP_PORT", 8080),
			},
			Database: Mysql{
				Host:     getString("DB_HOST", "localhost"),
				Port:     getInt("DB_PORT", 3306),
				Username: getString("DB_USER", "root"),
				Password: getString("DB_PASS", "password"),
				Name:     getString("DB_NAME", "myapp"),
			},
		}

		log.Println("✅ Config loaded")
	})

}

func getString(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func getInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalln("string to int conversion failed for key :", key, " error : ", err)
	}

	return i
}

func getBool(key string, def bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Fatalln("string to bool conversion failed for key :", key, " error : ", err)
	}

	return b
}
