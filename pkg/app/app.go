package app

import (
	"FreeMusic/pkg/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	handler "FreeMusic/pkg/delivery/http/v1"
	"FreeMusic/pkg/repository"
	"FreeMusic/pkg/service"
)

func Run(configPath string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//if err := initConfig(); err != nil {
	//	logrus.Fatalf("error initializing configs: %s", err.Error())
	//}
	//
	//if err := godotenv.Load(); err != nil {
	//	logrus.Fatalf("error loading env variables: %s", err.Error())
	//}

	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//})
	//if err != nil {
	//	logrus.Fatalf("failed to initialize db: %s", err.Error())
	//}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		//if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		//	logrus.Fatalf("error occured while running http server: %s", err.Error())
		//}
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	//
	//if err := db.Close(); err != nil {
	//	logrus.Errorf("error occured on db connection close: %s", err.Error())
	//}
}
