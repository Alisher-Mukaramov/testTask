package service

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"interviewTask/internal/models"
	"interviewTask/internal/repository"
)

var (
	ErrorEmpty   = errors.New("empty")
	ErrorProcess = errors.New("updating")
)

type Service interface {
	GetSdnList() (resp []byte, err error)
	Update(list models.SdnList) (err error)
	State() (int, error)
	GetNames(names []string, _type string) ([]models.Name, error)
}

type service struct {
	repo    repository.Repository
	httpLib *resty.Client
}

func New(options ...func(r *service)) Service {
	svc := &service{}
	for _, opt := range options {
		opt(svc)
	}
	return svc
}
