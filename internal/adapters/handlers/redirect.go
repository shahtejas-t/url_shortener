package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shahtejas-t/url_shortener/internal/core/services"
)

type RedirectFunctionHandler struct {
	linkService *services.LinkService
}

func NewRedirectFunctionHandler(l *services.LinkService) *RedirectFunctionHandler {
	return &RedirectFunctionHandler{linkService: l}
}

func (h *RedirectFunctionHandler) Redirect(c *gin.Context) {
	ctx := c.Request.Context()
	pathSegments := strings.Split(c.Request.URL.Path, "/")
	if len(pathSegments) < 2 {
		ClientError(c.Writer, http.StatusBadRequest, "Invalid URL path")
		return
	}

	shortLinkKey := pathSegments[len(pathSegments)-1]
	longLink, err := h.linkService.GetOriginalURL(ctx, shortLinkKey)
	if err != nil || *longLink == "" {
		ClientError(c.Writer, http.StatusNotFound, "Link not found")
		return
	}

	http.Redirect(c.Writer, c.Request, *longLink, http.StatusMovedPermanently)
}
