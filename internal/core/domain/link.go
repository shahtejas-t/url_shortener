package domain

type Link struct {
	Id          string `bson:"_id" json:"id"`
	OriginalURL string `bson:"original_url" json:"original_url"`
}
