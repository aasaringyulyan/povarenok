package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"scraping"
	"scraping/pkg/handler"
	"scraping/pkg/service"
	"syscall"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(scraping.Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("scraping Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("scraping Shutting Down")
}
