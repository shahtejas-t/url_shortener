package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	mongoCollectionName string
	mongoDatabaseName   string
	mongoAddress        string
	redisAddress        string
	redisPassword       string
	redisDB             int
}

func NewConfig() *AppConfig {
	return &AppConfig{
		mongoCollectionName: "Url_shortener_collection",
		mongoDatabaseName:   "Url_shortener",
		mongoAddress:        "mongodb://localhost:27017",
		redisAddress:        "redis://localhost:6379",
		redisPassword:       "",
		redisDB:             0,
	}
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file: ", err)
	}
}

func (c *AppConfig) GetMongoParams() (string, string, string) {
	mongoAddress, addrOk := os.LookupEnv("mongoAddress")
	collectionName, collectionOk := os.LookupEnv("mongoCollectionName")
	dbName, dbOk := os.LookupEnv("mongoDatabaseName")

	if !addrOk {
		fmt.Println("Need mongoAddress environment variable")
		return c.mongoAddress, c.mongoDatabaseName, c.mongoCollectionName
	}

	if !dbOk {
		fmt.Println("Need mongodb db environment variable")
		return mongoAddress, c.mongoDatabaseName, c.mongoCollectionName
	}

	if !collectionOk {
		fmt.Println("Need mongodb collection environment variable")
		return mongoAddress, dbName, c.mongoCollectionName
	}
	return mongoAddress, dbName, collectionName
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
