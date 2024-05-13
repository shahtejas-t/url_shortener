package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	linkCollectionName   string
	metricCollectionName string
	mongoDatabaseName    string
	mongoAddress         string
	redisAddress         string
	redisPassword        string
	redisDB              int
}

func NewConfig() *AppConfig {
	return &AppConfig{
		linkCollectionName:   "link",
		metricCollectionName: "metric",
		mongoDatabaseName:    "Url_shortener",
		mongoAddress:         "mongodb://mongodb:27017",
		redisAddress:         "redis://redis:6379",
		redisPassword:        "",
		redisDB:              0,
	}
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file: ", err)
	}
}

func (c *AppConfig) GetMongoParams() (string, string, string, string) {
	mongoAddress, addrOk := os.LookupEnv("mongoAddress")
	linkCollectionName, linkCollectionOk := os.LookupEnv("linkCollectionName")
	metricCollectionName, metricCollectionOk := os.LookupEnv("metricCollectionName")
	dbName, dbOk := os.LookupEnv("mongoDatabaseName")

	if !addrOk {
		fmt.Println("Need mongoAddress environment variable")
		return c.mongoAddress, c.mongoDatabaseName, c.linkCollectionName, c.metricCollectionName
	}

	if !dbOk {
		fmt.Println("Need mongodb db environment variable")
		return mongoAddress, c.mongoDatabaseName, c.linkCollectionName, c.metricCollectionName
	}

	if !linkCollectionOk || !metricCollectionOk {
		fmt.Println("Need mongodb collection environment variable")
		return mongoAddress, dbName, c.linkCollectionName, c.metricCollectionName
	}
	return mongoAddress, dbName, linkCollectionName, metricCollectionName
}

func (c *AppConfig) GetRedisParams() (string, string, int) {
	address, ok := os.LookupEnv("RedisAddress")
	if !ok {
		fmt.Println("Need RedisAddress environment variable")
		return c.redisAddress, c.redisPassword, c.redisDB
	}

	password, ok := os.LookupEnv("RedisPassword")
	if !ok {
		fmt.Println("Need RedisPassword environment variable")
		return address, c.redisPassword, c.redisDB
	}

	dbStr, ok := os.LookupEnv("RedisDB")
	if !ok {
		fmt.Println("Need RedisDB environment variable")
		return address, password, c.redisDB
	}

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		fmt.Printf("RedisDB environment variable is not a valid integer: %v\n", err)
		return address, password, c.redisDB
	}

	return address, password, db
}
