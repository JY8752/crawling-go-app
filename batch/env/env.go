package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	var err error
	if os.Getenv("ENVIROMENT") == "prod" {
		err = godotenv.Load("../../.env.prod")
	} else {
		err = godotenv.Load("../../.env")
	}

	if err != nil {
		log.Fatalf("not found env file. err: %v\n", err.Error())
	}
}
