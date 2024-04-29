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
	mongoAddress, mongoDbName, mongoCollectionName := appConfig.GetMongoParams()
	ctx := context.TODO()

	cache := cache.NewRedisCache(redisAddress, redisPassword, redisDB)

	// linkRepo := repository.NewLinkRepository(context.TODO(), linkTableName)
	linkRepo, err := repository.NewLinkRepository(ctx, mongoAddress, mongoDbName, mongoCollectionName)
	if err != nil {
		fmt.Printf("Error creating link repository: %s\n", err)
		return
	}

	linkService := services.NewLinkService(linkRepo, cache)

	deleteHandler := handlers.NewDeleteFunctionHandler(linkService)
	generateHandler := handlers.NewGenerateLinkFunctionHandler(linkService)
	redirectHandler := handlers.NewRedirectFunctionHandler(linkService)

	router := gin.Default()

	// routes
	router.DELETE("/shorten", deleteHandler.Delete)
	router.POST("/shorten", generateHandler.CreateShortLink)
	router.GET("/:shortLink", redirectHandler.Redirect)

	router.Run(":8080")
}
