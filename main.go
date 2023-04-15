package main

import (
	"mm-api/database"
	"mm-api/jwt"
	"os"
	"time"

	"mm-api/common"

	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)
func main() {

	err := database.InitMongoDb()
	if err != nil {
		log.Println(err)
		panic(err)
	}


	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	common.Rdb = redis.NewClient(opt)
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	app.Use(limiter.New(limiter.Config{
		Max: 1,
		Expiration: 1 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{"success": false, "message": "Too many requests"})
		},
		SkipFailedRequests: true,
	},))

	app.Post("/api/v1/login", jwt.JWTHandler)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://opensea.io",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/api/v1/data/:contract", database.GetIndexedContractHandler)
	
	app.Listen(":" + port)

}