package domain

import "time"

type Link struct {
	Id          string    `bson:"_id" json:"id"`
	OriginalURL string    `bson:"original_url" json:"original_url"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}
