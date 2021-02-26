package main

import (
	"goapi_admin_products/infrastructure"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	//	"github.com/ivan-salazar14/apigo_products/infrastructure"
)

func main() {
	log.Println("starting API cmd")
	port := os.Getenv("API_PORT")
	infrastructure.Start(port)
}
