package app

import (
	"FreeMusic/internal/config"
	"FreeMusic/internal/repository/mongodb"
	"FreeMusic/internal/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	handler "FreeMusic/internal/delivery/http/v1"
	"FreeMusic/internal/repository"
	"FreeMusic/internal/service"
)

func Run(configPath string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %v", err)
	}

	fileStorage, err := mongodb.NewMongoFileStorage(*config)
	if err != nil {
		logrus.Fatalf("error initializing configs: %v", err)
	}

	repos := repository.NewRepository(fileStorage)

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer(config)
	go func() {
		if err := srv.Run(handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("FreeMusic Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("FreeMusic Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	//
	//if err := db.Close(); err != nil {
	//	logrus.Errorf("error occured on db connection close: %s", err.Error())
	//}
}
