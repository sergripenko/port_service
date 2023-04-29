package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/sergripenko/port_service/internal/service/port"
	"github.com/sergripenko/port_service/internal/service/port/repository/mem"
)

type App struct {
	portService port.ServiceProvider
}

func NewApp() (*App, error) {
	repo := mem.NewRepository()
	portService := port.NewService(repo)

	return &App{
		portService: portService,
	}, nil
}

func (a *App) Run() error {
	fileName := "ports.json"
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close() //nolint

	ctx := context.TODO()

	// Create done chan for graceful shutdown.
	doneChan := make(chan struct{})

	go func() {
		err = a.portService.AddPorts(ctx, f)
		doneChan <- struct{}{}
		if err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) //nolint

	select {
	case <-quit:
		log.Print("shutdown")
		a.portService.Shutdown()
	case <-doneChan:
	}
	return err
}
