package ports

import (
	"context"

	"github.com/shahtejas-t/url_shortener/internal/core/domain"
)

type MetricPort interface {
	All(context.Context) ([]domain.Metric, error)
	Get(context.Context, string) (domain.Metric, error)
	GetShortLinkRecordCount(context.Context, string) (int64, error)
	GetTopShortLinksByRecordCount(context.Context, int64) ([]domain.LinkMetric, error)
	Create(context.Context, domain.Metric) error
	DeleteByShortLink(context.Context, string) error
}
