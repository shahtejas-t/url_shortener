package domain

import "time"

type Metric struct {
	ShortLink   string    `bson:"short_link" json:"short_link"`
	OriginalURL string    `bson:"original_url" json:"original_url"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}

type LinkMetric struct {
	OriginalURL string `bson:"original_url" json:"original_url"`
	Count       int    `bson:"count" json:"count"`
}
