package repository

import (
	"context"
	"fmt"

	"github.com/shahtejas-t/url_shortener/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LinkRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewLinkRepository(ctx context.Context, uri, dbName, collectionName string) (*LinkRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB server: %w", err)
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &LinkRepository{
		client:     client,
		collection: collection,
	}, nil
}

func (r *LinkRepository) All(ctx context.Context) ([]domain.Link, error) {
	var links []domain.Link

	cursor, err := r.collection.Find(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents in MongoDB: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var link domain.Link
		if err := cursor.Decode(&link); err != nil {
			return nil, fmt.Errorf("failed to decode document: %w", err)
		}
		links = append(links, link)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return links, nil
}

func (r *LinkRepository) Get(ctx context.Context, id string) (domain.Link, error) {
	var link domain.Link

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&link)
	if err != nil {
		return domain.Link{}, fmt.Errorf("failed to find document with ID %s: %w", id, err)
	}

	return link, nil
}

func (r *LinkRepository) Create(ctx context.Context, link domain.Link) error {
	_, err := r.collection.InsertOne(ctx, link)
	if err != nil {
		return fmt.Errorf("failed to insert document into MongoDB: %w", err)
	}
	return nil
}

func (r *LinkRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete document with ID %s: %w", id, err)
	}
	return nil
}
