package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"moul.io/zapgorm2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/ivas1ly/waybill-app/database"
	"github.com/ivas1ly/waybill-app/domain"
	"github.com/ivas1ly/waybill-app/graph"
	"github.com/ivas1ly/waybill-app/graph/generated"
	"github.com/ivas1ly/waybill-app/models"
)

type App struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewApp() *App {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db := database.DBConn

	return &App{
		logger: logger,
		db:     db,
	}
}

func (a *App) Run(port string) {
	a.initDatabase()

	usersRepository := database.UsersRepository{DB: database.DBConn}

	d := domain.NewDomain(usersRepository)

	app := fiber.New(fiber.Config{
		ServerHeader:          "Fiber",
		BodyLimit:             4 * 1024 * 1024,
		Concurrency:           256 * 1024,
		ReadTimeout:           2 * time.Minute,
		WriteTimeout:          2 * time.Minute,
		DisableStartupMessage: true,
	})
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{Domain: d},
	}))
	srv.Use(extension.FixedComplexityLimit(300))
	gqlHandler := srv.Handler()
	pg := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		pg(c.Context())
		return nil
	})

	a.logger.Info(fmt.Sprintf("Connect to http://localhost:%s/ for GraphQL playground", port))

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":" + port); err != nil {
			a.logger.Fatal(fmt.Sprintf("%+v", err))
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	a.logger.Info("Gracefully shutting down...")
	_ = app.Shutdown()

	a.logger.Info("Running cleanup tasks...")
	// Your cleanup tasks go here
	a.logger.Info("Fiber was successful shutdown.")
}

func (a *App) initDatabase() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.dbname"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.timezone"))

	gormLogger := zapgorm2.New(a.logger)
	gormLogger.SetAsDefault()

	a.db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	}), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		a.logger.Fatal("Failed to Connect database!")
	}
	a.logger.Info("Connection Opened to Database")

	err = a.db.AutoMigrate(&models.Car{}, &models.Driver{}, &models.Waybill{}, &models.User{})
	if err != nil {
		a.logger.Fatal("Database Not Migrated")
	}
	a.logger.Info("Database Migrated")
}
