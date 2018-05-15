package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//NewInstance("ami-43a15f3e", "t2.micro", "Tester T. Testerson", 1, 1)
	GetAmi()
}

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Print("Failed to load local environment file")
	}
}
