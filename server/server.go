package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ivas1ly/waybill-app/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/gofiber/fiber/v2"

	"github.com/spf13/viper"

	"go.uber.org/zap"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"

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
	zapLogger, _ := zap.NewProduction()
	defer func(zapLogger *zap.Logger) {
		_ = zapLogger.Sync()
	}(zapLogger)
	db := ConnectDatabase(zapLogger)

	return &App{
		logger: zapLogger,
		db:     db,
	}
}

func (a *App) Run(port string) {

	usersRepository := database.UsersRepository{DB: a.db}
	waybillsRepository := database.WaybillsRepository{DB: a.db}
	driversRepository := database.DriversRepository{DB: a.db}
	carsRepository := database.CarsRepository{DB: a.db}

	d := domain.NewDomain(a.logger, usersRepository, waybillsRepository, driversRepository, carsRepository)

	app := fiber.New(fiber.Config{
		ServerHeader:          "Fiber",
		BodyLimit:             4 * 1024 * 1024,
		Concurrency:           256 * 1024,
		ReadTimeout:           2 * time.Minute,
		WriteTimeout:          2 * time.Minute,
		DisableStartupMessage: false,
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
		CookieHTTPOnly: true,
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
	app.Use(middleware.Auth(d))

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

func ConnectDatabase(logger *zap.Logger) *gorm.DB {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.dbname"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.timezone"))

	newLogger := gl.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gl.Config{
		SlowThreshold:             time.Second,
		Colorful:                  true,
		IgnoreRecordNotFoundError: false,
		LogLevel:                  gl.Info,
	})

	db := database.DB
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		logger.Fatal("Failed to Connect database!")
	}
	logger.Info("Connection Opened to Database")

	err = db.AutoMigrate(&models.Car{}, &models.Driver{}, &models.Waybill{}, &models.User{})

	if err != nil {
		logger.Fatal("Database Not Migrated")
	}
	logger.Info("Database Migrated")

	return db
}
