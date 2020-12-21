package main

import (
	"log"

	"github.com/NODO-UH/uh-email-quota/src/api"
)

func main() {
	// Start API server
	if err := api.StartAPI(); err != nil {
		log.Println(err.Error())
	}
}
