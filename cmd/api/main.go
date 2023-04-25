package main

import (
	"context"
	"github.com/dozheiny/backend-template/database"
	"github.com/dozheiny/backend-template/pkg/object"
	"github.com/dozheiny/backend-template/pkg/redis"
	"github.com/gofiber/fiber/v2"
	"log"
)

// init function will run before running the main function.
// that helps to health check your environment is healthy or not.
func init() {

	ctx := context.Background()

	if err := database.Connection(ctx); err != nil {
		log.Fatalf("Gotcha Error on make connection to mongo: %s", err.Error())
	}

	if _, err := redis.GetRedis().GetConnection(); err != nil {
		log.Fatalf("Gotcha Error on get connection from redis: %s", err.Error())
	}

	if err := object.Initialize(); err != nil {
		log.Fatalf("Gotcha Error on make connection to minio: %s", err.Error())
	}
}

func main() {
	route := fiber.New()

	if err := route.Listen(":8080"); err != nil {
		log.Fatalf("Got ERR on listening service: %s", err.Error())
	}
}
