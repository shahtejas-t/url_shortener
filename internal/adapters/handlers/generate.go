package handlers

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shahtejas-t/url_shortener/internal/core/domain"
	"github.com/shahtejas-t/url_shortener/internal/core/services"
)

type RequestBody struct {
	Long string `json:"long"`
}

type GenerateLinkFunctionHandler struct {
	linkService   *services.LinkService
	metricService *services.MetricService
}

func NewGenerateLinkFunctionHandler(l *services.LinkService, m *services.MetricService) *GenerateLinkFunctionHandler {
	return &GenerateLinkFunctionHandler{linkService: l, metricService: m}
}

func (h *GenerateLinkFunctionHandler) CreateShortLink(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody RequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&requestBody)
	if err != nil {
		ClientError(c.Writer, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if requestBody.Long == "" {
		ClientError(c.Writer, http.StatusBadRequest, "URL cannot be empty")
		return
	}
	if len(requestBody.Long) < 15 {
		ClientError(c.Writer, http.StatusBadRequest, "URL must be at least 15 characters long")
		return
	}
	if !IsValidLink(requestBody.Long) {
		ClientError(c.Writer, http.StatusBadRequest, "Invalid URL format")
		return
	}

	link := domain.Link{
		Id:          GenerateShortURLID(8),
		OriginalURL: requestBody.Long,
	}

	link, err = h.linkService.Create(ctx, link)
	if err != nil {
		ServerError(c.Writer, err)
		return
	}

	js, err := json.Marshal(link)
	if err != nil {
		ServerError(c.Writer, err)
		return
	}

	metric := domain.Metric{
		ShortLink:   link.Id,
		OriginalURL: link.OriginalURL,
		CreatedAt:   time.Now(),
	}

	err = h.metricService.Create(ctx, metric)
	if err != nil {
		ServerError(c.Writer, err)
		return
	}

	log.Printf("The system generated a short URL with the ID %s\n", link.Id)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(js)
}

func GenerateShortURLID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[charIndex.Int64()]
	}
	return string(result)
}
