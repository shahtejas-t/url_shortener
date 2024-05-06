package repository

import (
	"context"
	"fmt"

	"github.com/shahtejas-t/url_shortener/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MetricRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMetricRepository(ctx context.Context, uri, dbName, collectionName string) (*MetricRepository, error) {
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

	return &MetricRepository{
		client:     client,
		collection: collection,
	}, nil
}

func (r *MetricRepository) All(ctx context.Context) ([]domain.Metric, error) {
	var metrics []domain.Metric

	cursor, err := r.collection.Find(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents in MongoDB: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var metric domain.Metric
		if err := cursor.Decode(&metric); err != nil {
			return nil, fmt.Errorf("failed to decode document: %w", err)
		}
		metrics = append(metrics, metric)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return metrics, nil
}

func (r *MetricRepository) Get(ctx context.Context, id string) (domain.Metric, error) {
	var metric domain.Metric

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&metric)
	if err != nil {
		return domain.Metric{}, fmt.Errorf("failed to find document with ID %s: %w", id, err)
	}

	return metric, nil
}

func (r *MetricRepository) GetShortLinkRecordCount(ctx context.Context, shortLink string) (int64, error) {
	filter := bson.M{"short_link": shortLink}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *MetricRepository) GetTopShortLinksByRecordCount(ctx context.Context, limit int64) ([]domain.LinkMetric, error) {

	linkMetrics := []domain.LinkMetric{}

	pipeline := bson.A{
		bson.M{"$group": bson.M{
			"_id":          "$short_link",
			"original_url": bson.M{"$first": "$original_url"},
			"count":        bson.M{"$sum": 1},
		}},
		bson.M{"$sort": bson.M{"count": -1}},
		bson.M{"$limit": limit},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &linkMetrics); err != nil {
		return nil, err
	}

	return linkMetrics, nil
}

func (r *MetricRepository) Create(ctx context.Context, metric domain.Metric) error {
	_, err := r.collection.InsertOne(ctx, metric)
	if err != nil {
		return fmt.Errorf("failed to insert document into MongoDB: %w", err)
	}
	return nil
}

func (r *MetricRepository) DeleteByShortLink(ctx context.Context, shortLink string) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"short_link": shortLink})
	if err != nil {
		return fmt.Errorf("failed to delete documents with short link %s: %w", shortLink, err)
	}
	return nil
}
