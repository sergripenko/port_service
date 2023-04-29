package main

import (
	"log"

	"github.com/sergripenko/port_service/internal/app"
)

func main() {
	portsApp, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	if err := portsApp.Run(); err != nil {
		log.Fatal(err)
	}
}
