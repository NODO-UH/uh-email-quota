package main

import (
	"log"

	"github.com/NODO-UH/uh-email-quota/src/api"
	grpc "github.com/NODO-UH/uh-email-quota/src/grpc"
)

func main() {
	// Start gRPC
	go grpc.StartGRPC()
	// Start API server
	if err := api.StartAPI(); err != nil {
		log.Println(err.Error())
	}
}
