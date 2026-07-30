package config 

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(){

	err := godotenv.Load()

	if(err != nil){
		log.Fatal("Error loading .env file");
	}

	log.Println("Enviroment variables loaded successfully")
}