package main

import (
	"os"
	"os/signal"
	"scraping"
	"scraping/pkg/handler"
	"scraping/pkg/logging"
	"scraping/pkg/service"
	"syscall"
)

func main() {
	logger := logging.GetLogger()
	services := service.NewService(logger)
	handlers := handler.NewHandler(logger, services)

	srv := new(scraping.Server)
	go func() {
		if err := srv.Run("1234", handlers.InitRoutes()); err != nil {
			logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Info("scraping Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("scraping Shutting Down")
}
