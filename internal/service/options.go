package service

import (
	"github.com/go-resty/resty/v2"
	"interviewTask/internal/repository"
)

func WithRepository(r repository.Repository) func(*service) {
	return func(s *service) {
		s.repo = r
	}
}
func WithHttpLib(c *resty.Client) func(*service) {
	return func(s *service) {
		s.httpLib = c
	}
}
