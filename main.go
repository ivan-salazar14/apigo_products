package main

import (
	"log"

	"github.com/ivan-salazar14/apigo_products/infrastructure"
)

func main() {
	log.Println("starting API cmd")
	port := "8081" // os.Getenv("API_PORT")
	infrastructure.Start(port)
}
