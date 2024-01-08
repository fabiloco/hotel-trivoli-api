package config

import (
	"os"
	"github.com/joho/godotenv"
)

func Config(key string) string {
  // load env file
  err := godotenv.Load(".env")

	if err != nil {
		//fmt.Print("Error loading .env file. Please, check again if you have the file.")
		panic("Error loading .env file. Please, check again if you have the file.")
	}

	return os.Getenv(key)
}
