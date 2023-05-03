package handler

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"interviewTask/internal/models"
	"log"
	"net/http"
)

func (h *Handler) Update(c *gin.Context) {
	var (
		err     error
		resp    = models.Response{}
		sdnList = models.SdnList{}
	)

	defer func() {
		switch {
		case err != nil:
			resp.SetStatus(false, "service unavailable", http.StatusServiceUnavailable)
			log.Printf("%s", err.Error())
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, resp)
		default:
			resp.SetStatus(true, "", http.StatusOK)
			c.JSON(http.StatusOK, resp)
		}
	}()

	result, err := h.svc.GetSdnList()
	if err != nil {
		return
	}

	if err = xml.Unmarshal(result, &sdnList); err != nil {
		return
	}

	err = h.svc.Update(sdnList)
}
