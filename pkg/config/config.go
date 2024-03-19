package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvironmet() error {
	err := godotenv.Load("../../.env")
	env := os.Getenv("ENV")
	if err != nil && env == "" {
		log.Fatalf("failed to load env, err : %v", err)
	}
	return nil
}

// func LoadEnvironmet() error {
// 	if err := godotenv.Load("../../.env"); err != nil {
// 		log.Fatalf("failed to load env, err : %v", err)
// 	}
// 	return nil
// }
