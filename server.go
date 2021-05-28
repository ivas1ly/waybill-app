package main

import (
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"go.uber.org/zap"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/ivas1ly/waybill-app/graph"
	"github.com/ivas1ly/waybill-app/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET,POST",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "cookie:csrf_waybill",
		CookieName:     "csrf_waybill",
		CookieHTTPOnly: false,
		CookieSameSite: "Strict",
		Expiration:     30 * time.Minute,
		KeyGenerator:   utils.UUIDv4,
		ContextKey:     "csrf_waybill",
	}))

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		Max:        20,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(extension.FixedComplexityLimit(300))
	gqlHandler := srv.Handler()
	pg := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	sugar.Info(time.Now())

	app.All("/", func(c *fiber.Ctx) error {
		pg(c.Context())
		return nil
	})

	sugar.Infof("Connect to http://localhost:%s/ for GraphQL playground", port)
	sugar.Fatal(app.Listen(":" + port))
}
