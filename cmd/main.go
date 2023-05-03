package main

import (
	"github.com/go-resty/resty/v2"
	"interviewTask/internal/handler"
	"interviewTask/internal/repository"
	"interviewTask/internal/server"
	"interviewTask/internal/service"
	"interviewTask/pkg/database/driver/sqlx/postgres"
)

func main() {

	// hardcore
	pgConfig := &postgres.Config{
		Host:     "postgres",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "sdn",
	}

	// инициализация БД
	pg := postgres.New(pgConfig)

	// инициализируем слоя репозитория
	repo := repository.New(repository.WithPostgresService(pg))

	// инициализируем сервиса
	svc := service.New(
		service.WithRepository(repo),
		service.WithHttpLib(resty.New()),
	)

	// инициализация обработчика
	h := handler.New(handler.WithService(svc))

	// инициализация сервера
	srv := server.New(h)

	// запуск сервера
	srv.Run()
}
