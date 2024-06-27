package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lavatee/messenger"
	"github.com/lavatee/messenger/internal/endpoint"
	"github.com/lavatee/messenger/internal/repository"
	"github.com/lavatee/messenger/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		log.Fatalf("config err", err.Error())
	}
	wd, err := os.Getwd()
	if err := godotenv.Load(wd + "/.env"); err != nil {
		log.Fatalf("env err", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{Port: viper.GetString("db.port"), Host: viper.GetString("db.host"), Username: viper.GetString("db.username"), Password: os.Getenv("DB_PASSWORD"), DBName: viper.GetString("db.dbname"), SSLmode: viper.GetString("db.sslmode")})
	if err != nil {
		log.Fatalf("db err", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := endpoint.NewEndpoint(services)
	srv := new(messenger.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf(err.Error())
	}
}
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
