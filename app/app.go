package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/electro98/noteapp/domain"
	"github.com/electro98/noteapp/models"
	"github.com/electro98/noteapp/services/note"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() {
	initConfiguration()

	logger := initLogger()

	db := initDatabase(logger)

	servicesLogger := logger.With().
		Str("component", "services_internals").
		Logger()
	noteUseCase := initServices(db, &servicesLogger)

	// Set up middleware
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			logger.Err(v.Error).
				Str("uri", v.URI).
				Int("status", v.Status).
				Str("method", v.Method).
				Err(v.Error).
				Msg("request")
			return nil
		},
	}))
	e.Use(middleware.CORS("*")) // Security RISK
	e.Use(middleware.ContextTimeout(time.Second))

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	// Register services
	servicesLogger = logger.With().
		Str("component", "services").
		Logger()
	api := e.Group("/api")
	note.NewNoteHandler(api, noteUseCase, &servicesLogger)

	// Graceful Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	port := viper.GetInt("app.port")
	sc := echo.StartConfig{
		Address:         fmt.Sprintf(":%d", port),
		GracefulTimeout: 3 * time.Second,
	}
	if err := sc.Start(ctx, e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func initConfiguration() {
	viper.SetConfigType("toml")

	viper.AddConfigPath(".")
	viper.SetConfigName(".config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}

func initLogger() *zerolog.Logger {
	var writer io.Writer
	if viper.GetBool("logger.prettyLogging") {
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	} else {
		writer = os.Stderr
	}
	logger := zerolog.New(writer).
		With().Timestamp().Logger()

	return &logger
}

func initDatabase(logger *zerolog.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo",
		viper.GetString("database.host"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetInt("database.port"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect database")
	}

	db.AutoMigrate(&models.Note{})

	logger.Info().Msg("started database")
	return db
}

func initServices(db *gorm.DB, logger *zerolog.Logger) domain.NoteUseCase {
	noteRepository := note.NewNoteRepository(db)

	noteUseCase := note.NewNoteUseCase(noteRepository, logger)

	return noteUseCase
}
