package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"

	"github.com/spf13/viper"

	"go.uber.org/zap"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"

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

	d := domain.NewDomain(usersRepository, waybillsRepository, driversRepository, carsRepository)

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
		//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug: false,
		//MaxAge:           300,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//router.Use(appCustomMiddleware.Auth(usersRepository))
	/*router.Use(csrf.Protect(
		[]byte("a-32-byte-long-key-goes-here"),
		csrf.TrustedOrigins([]string{"*"}),
		csrf.Secure(false),
		csrf.HttpOnly(false),
	))*/

	router.Route("/query", func(r chi.Router) {
		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &graph.Resolver{Domain: d}})
		srv := handler.NewDefaultServer(schema)
		srv.Use(extension.FixedComplexityLimit(300))

		r.Handle("/", srv)
	})

	pg := playground.Handler("GraphQL playground", "/query")
	router.Get("/", pg)

	a.logger.Info(fmt.Sprintf("Connect to http://localhost:%s/ for GraphQL playground", port))
	if err := http.ListenAndServe(":"+port, router); err != nil {
		a.logger.Fatal(fmt.Sprintf("%+v", err))
	}
	/*
		// Listen from a different goroutine
		go func() {
			if err := http.ListenAndServe(":"+port, router); err != nil {
				a.logger.Fatal(fmt.Sprintf("%+v", err))
			}
		}()

		c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
		signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

		_ = <-c // This blocks the main thread until an interrupt is received
		a.logger.Info("Gracefully shutting down...")
		//_ = http.Server.Shutdown(router, context.Background())

		a.logger.Info("Running cleanup tasks...")
		// Your cleanup tasks go here
		a.logger.Info("Server was successful shutdown."*/
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
