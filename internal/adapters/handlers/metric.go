package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahtejas-t/url_shortener/internal/core/services"
)

type MetricHandler struct {
	metricService *services.MetricService
}

func NewMetricHandler(metricService *services.MetricService) *MetricHandler {
	return &MetricHandler{metricService: metricService}
}

func (h *MetricHandler) GetTopShortLinksByRecordCount(c *gin.Context) {
	ctx := c.Request.Context()
	var requestBody struct {
		Limit int64 `json:"limit"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		ServerError(c.Writer, err)
		return
	}

	result, err := h.metricService.GetTopShortLinksByRecordCount(ctx, requestBody.Limit)
	if err != nil {
		ServerError(c.Writer, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *MetricHandler) GetShortLinkRecordCount(c *gin.Context) {
	ctx := c.Request.Context()
	var requestBody struct {
		ShortLink string `json:"short_link"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		ServerError(c.Writer, err)
		return
	}

	count, err := h.metricService.GetShortLinkRecordCount(ctx, requestBody.ShortLink)
	if err != nil {
		ServerError(c.Writer, err)
		return
	}
	response := map[string]interface{}{
		"short_link": requestBody.ShortLink,
		"count":      count,
	}
	c.JSON(http.StatusOK, response)
}
