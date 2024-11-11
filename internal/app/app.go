package app

import (
	"log"
	"auth-service-backend/config"
	"auth-service-backend/internal/database/postgres"
	"auth-service-backend/internal/database/redis"
	"auth-service-backend/pkg/minio"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializationApp() error {
	err := config.InitConfig()
	if err != nil {
		return err
	}

	database := postgres.NewDataBase()
	err = database.InitializationDB()
	if err != nil {
		return err
	}

	redis := redis.NewRedis(config.AppConfig.RedisHost)

	minio.Initialization(redis)

	app := fiber.New()
	app.Use(logger.New())

	diContainer := NewDIContainer(database.GetQueries(), app, redis)
	diContainer.InitializationModules()

	log.Fatal(app.Listen(":" + config.AppConfig.Port))

	return nil
}
