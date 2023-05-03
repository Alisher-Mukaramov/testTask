package handler

import (
	"github.com/gin-gonic/gin"
	"interviewTask/internal/models"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) GetNames(c *gin.Context) {
	var (
		err   error
		names []models.Name
		resp  = models.Response{}
	)

	_name := c.Query("name")
	_type := c.Query("type")

	_names := strings.Split(_name, " ")

	defer func() {
		switch {
		case err != nil:
			resp.SetStatus(false, "service unavailable", http.StatusServiceUnavailable)
			log.Printf("%s", err.Error())
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, resp)
		default:
			c.JSON(http.StatusOK, names)
		}
	}()

	names, err = h.svc.GetNames(_names, _type)

	if len(names) == 0 {
		names = []models.Name{}
	}
}
