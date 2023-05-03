package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"interviewTask/internal/models"
	"interviewTask/internal/service"
	"log"
	"net/http"
)

func (h *Handler) State(c *gin.Context) {
	var (
		err  error
		resp = models.Response{}
	)

	defer func() {
		switch {
		case errors.Is(err, service.ErrorProcess):
			resp.SetStatus(false, service.ErrorProcess.Error(), 0)
			c.JSON(http.StatusServiceUnavailable, resp)
		case errors.Is(err, service.ErrorEmpty):
			resp.SetStatus(false, service.ErrorEmpty.Error(), 0)
			c.JSON(http.StatusServiceUnavailable, resp)
		case err != nil:
			resp.SetStatus(false, "service unavailable", http.StatusServiceUnavailable)
			log.Printf("%s", err.Error())
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, resp)
		default:
			resp.SetStatus(true, "ok", 0)
			c.JSON(http.StatusOK, resp)
		}
	}()

	count, err := h.svc.State()
	if err != nil {
		return
	}

	if count == 0 {
		err = service.ErrorEmpty
	}
}
