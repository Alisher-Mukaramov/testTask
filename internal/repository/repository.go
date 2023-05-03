package repository

import (
	"interviewTask/internal/models"
	"interviewTask/pkg/database/driver/sqlx/postgres"
)

type Repository interface {
	Update(XMLString string) error
	GetNames(n []string, t string) (names []models.Name, err error)
	State() (count int, err error)
	IsLocked() bool
}

type repository struct {
	pg     *postgres.Postgres
	locked bool
}

func New(options ...func(r *repository)) Repository {
	repo := &repository{}
	for _, opt := range options {
		opt(repo)
	}
	return repo
}
