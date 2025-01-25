package server

import (
	config "api/config"
	"api/internal/application"
	"api/internal/application/handlers"
	"api/internal/infrastructure"
	"api/internal/infrastructure/init"
	serviceInit "api/internal/service/init"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoadConfig()

	cs := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", cfg.DBUser,
		cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db := infrastructure.MustInitDB(cs)
	defer db.Close()
	r := init.NewUnitOfWork(db)
	s := serviceInit.NewServices(*r, cfg.TokenConfig)
	handler := handlers.NewHandler(s)
	api := handler.InitRouters()
	srv := new(application.Server)
	go func() {
		if err := srv.Run(api, cfg); err != nil {
			panic(fmt.Sprintf("Error starting server: %v", err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Stop(context.Background()); err != nil {
		panic(fmt.Sprintf("Error stopping server: %v", err))
	}
}
