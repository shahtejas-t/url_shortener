// main.go
package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shahtejas-t/url_shortener/internal/adapters/cache"
	"github.com/shahtejas-t/url_shortener/internal/adapters/handlers"
	"github.com/shahtejas-t/url_shortener/internal/adapters/repository"
	"github.com/shahtejas-t/url_shortener/internal/config"
	"github.com/shahtejas-t/url_shortener/internal/core/services"
)

func main() {

	appConfig := config.NewConfig()
	redisAddress, redisPassword, redisDB := appConfig.GetRedisParams()
	mongoAddress, mongoDbName, linkCollectionName, metricCollectionName := appConfig.GetMongoParams()
	ctx := context.TODO()

	cache := cache.NewRedisCache(redisAddress, redisPassword, redisDB)

	linkRepo, err := repository.NewLinkRepository(ctx, mongoAddress, mongoDbName, linkCollectionName)
	if err != nil {
		fmt.Printf("Error creating link repository: %s\n", err)
		return
	}
	metricRepo, metricErr := repository.NewMetricRepository(ctx, mongoAddress, mongoDbName, metricCollectionName)

	if metricErr != nil {
		fmt.Printf("Error creating Metric repository: %s\n", metricErr)
		return
	}

	linkService := services.NewLinkService(linkRepo, cache)
	metricService := services.NewMetricService(metricRepo, cache)

	deleteHandler := handlers.NewDeleteFunctionHandler(linkService, metricService)
	generateHandler := handlers.NewGenerateLinkFunctionHandler(linkService, metricService)
	redirectHandler := handlers.NewRedirectFunctionHandler(linkService)

	metricHandler := handlers.NewMetricHandler(metricService)

	router := gin.Default()

	// routes
	router.DELETE("/shorten", deleteHandler.Delete)
	router.POST("/shorten", generateHandler.CreateShortLink)
	router.GET("/:shortLink", redirectHandler.Redirect)

	router.GET("/toprecords", metricHandler.GetTopShortLinksByRecordCount)
	router.GET("/linkcount", metricHandler.GetShortLinkRecordCount)

	router.Run(":8080")
}
