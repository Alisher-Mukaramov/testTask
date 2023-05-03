package handler

import (
	"interviewTask/internal/service"
)

func WithService(s service.Service) func(*Handler) {
	return func(h *Handler) {
		h.svc = s
	}
}
