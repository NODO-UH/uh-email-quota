package main

import (
	"flag"
	"log"

	"github.com/NODO-UH/uh-email-quota/src/api"
	conf "github.com/NODO-UH/uh-email-quota/src/config"
)

func main() {
	confPath := flag.String("conf", "config.json", "path to configuration file")
	flag.Parse()

	if err := conf.SetupConfiguration(*confPath); err != nil {
		panic(err)
	}

	// Start API server
	if err := api.StartAPI(); err != nil {
		log.Println(err.Error())
	}
}
