package app

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/sdf0106/ip-project/internal/delivery/http"
	"github.com/sdf0106/ip-project/internal/repository"
	"github.com/sdf0106/ip-project/internal/server"
	"github.com/sdf0106/ip-project/internal/service"
	"github.com/sdf0106/ip-project/pkg/auth"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error in intializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}

	tokenManager, err := auth.NewManager(os.Getenv("SIGNING_KEY"))

	if err != nil {
		logrus.Fatalf("error in initializting TokenManager: %s", err.Error())
	}

	config := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	dbPool := repository.NewPostgresDB(config)

	repos := repository.NewRepositories(dbPool)
	services := service.NewService(repos, tokenManager)
	handler := delivery.NewHandler(services)

	srv := server.NewServer(viper.GetString("port"), handler.InitRoutes())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 12 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logrus.Errorf("failed to stop server: %v", err)
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")

	return viper.ReadInConfig()
}
