package services

import (
	"context"
	"fmt"

	"github.com/shahtejas-t/url_shortener/internal/core/domain"
	"github.com/shahtejas-t/url_shortener/internal/core/ports"
)

type MetricService struct {
	port  ports.MetricPort
	cache ports.Cache
}

func NewMetricService(p ports.MetricPort, c ports.Cache) *MetricService {
	return &MetricService{port: p, cache: c}
}

func (service *MetricService) GetAll(ctx context.Context) ([]domain.Metric, error) {
	metrics, err := service.port.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all links: %w", err)
	}
	return metrics, nil
}

func (service *MetricService) GetShortLinkRecordCount(ctx context.Context, shortLinkKey string) (*int64, error) {
	count, err := service.port.GetShortLinkRecordCount(ctx, shortLinkKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get short URL for identifier '%s': %w", shortLinkKey, err)
	}
	return &count, nil
}

func (service *MetricService) GetTopShortLinksByRecordCount(ctx context.Context, limit int64) ([]domain.LinkMetric, error) {
	result, err := service.port.GetTopShortLinksByRecordCount(ctx, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get the details: %w", err)
	}
	return result, nil
}

func (service *MetricService) Create(ctx context.Context, metric domain.Metric) error {
	if service.port != nil {
		if err := service.port.Create(ctx, metric); err != nil {
			return fmt.Errorf("failed to create short URL: %w", err)
		}
		return nil
	}
	return fmt.Errorf("failed to create short URL : Server Error")
}

func (service *MetricService) Delete(ctx context.Context, shortLink string) error {
	if err := service.port.DeleteByShortLink(ctx, shortLink); err != nil {
		return fmt.Errorf("failed to delete short URL for identifier '%s': %w", shortLink, err)
	}
	return nil
}
