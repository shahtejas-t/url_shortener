package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahtejas-t/url_shortener/internal/core/services"
)

type DeleteFunctionHandler struct {
	linkService   *services.LinkService
	metricService *services.MetricService
}

func NewDeleteFunctionHandler(l *services.LinkService, m *services.MetricService) *DeleteFunctionHandler {
	return &DeleteFunctionHandler{linkService: l, metricService: m}
}

func (s *DeleteFunctionHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Request.URL.Query().Get("id")
	err := s.linkService.Delete(ctx, id)
	if err != nil {
		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = s.metricService.Delete(ctx, id)
	if err != nil {
		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}
