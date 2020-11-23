package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
)

// Config задает параметры конфигурации приложения
type Config struct {
	Port        string `envconfig:"PORT" default:"8080"`
	StaticsPath string `envconfig:"STATICS_PATH" default:"./static"`
}

func main() {
	config := new(Config)
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatalf("Can't process config: %v", err)
	}

	fs := http.FileServer(http.Dir(config.StaticsPath))
	http.Handle("/", fs)

	go func() {
		err = http.ListenAndServe(":"+config.Port, nil)
		if err != nil {
			log.Fatalf("Error while serving: %v", err)
		}
	}()

	log.Print("The app is ready")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	}

}
