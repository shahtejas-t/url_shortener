package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahtejas-t/url_shortener/internal/core/services"
)

type DeleteFunctionHandler struct {
	linkService *services.LinkService
}

func NewDeleteFunctionHandler(l *services.LinkService) *DeleteFunctionHandler {
	return &DeleteFunctionHandler{linkService: l}
}

func (s *DeleteFunctionHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Request.URL.Query().Get("id")
	err := s.linkService.Delete(ctx, id)
	if err != nil {
		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
