package main

import (
	"fmt"
	"log"
	"os"
	app "some_app/internal/api/http"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	addr, exists := os.LookupEnv("APP_ADDR")
	if !exists {
		log.Fatal("no addr")
	}
	port, exists := os.LookupEnv("APP_PORT")
	if !exists {
		log.Fatal("no port")
	}

	app.Run(fmt.Sprintf("%s:%s",addr, port))
}
