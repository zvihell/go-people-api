package main

import (
	"go-people-api/internal/controllers"
	"go-people-api/internal/repository"
	"go-people-api/internal/service"
	"go-people-api/pkg"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initLogrus() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})

	logLevel, err := log.ParseLevel(viper.GetString("logger.level"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error loading env variables: %s", err)
	}

	initLogrus()

	db, err := pkg.InitDB(pkg.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewDB(db)

	services := service.NewUserService(repo)

	handler := controllers.NewHandler(services)

	srv := &http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: handler.InitRouter(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("conf")
	return viper.ReadInConfig()
}
