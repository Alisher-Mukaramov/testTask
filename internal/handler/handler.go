package handler

import "interviewTask/internal/service"

type Handler struct {
	svc service.Service
}

func New(options ...func(h *Handler)) *Handler {
	h := &Handler{}
	for _, opt := range options {
		opt(h)
	}
	return h
}
