package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/electro98/noteapp/domain"
	"github.com/electro98/noteapp/models"
	"github.com/electro98/noteapp/services/note"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	db := createMockDatabase(&logger)

	servicesLogger := logger.With().
		Str("component", "services_internals").
		Logger()
	noteUseCase := initServices(db, &servicesLogger)

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

	servicesLogger = logger.With().
		Str("component", "services").
		Logger()
	api := e.Group("/api")
	note.NewNoteHandler(api, noteUseCase, &servicesLogger)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func createMockDatabase(logger *zerolog.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{TranslateError: true})
	if err != nil {
		logger.Error().Err(err).Msg("failed to connect database")
		panic("failed to connect database")
	}

	ctx := context.Background()

	db.AutoMigrate(&models.Note{})

	err = gorm.G[models.Note](db).Create(ctx, &models.Note{Title: "Note 1", Content: "Hello world!"})
	err = gorm.G[models.Note](db).Create(ctx, &models.Note{Title: "Note 2", Content: "This is a note."})

	logger.Info().Msg("started database")
	return db
}

func initServices(db *gorm.DB, logger *zerolog.Logger) domain.NoteUseCase {
	noteRepository := note.NewNoteRepository(db)

	noteUseCase := note.NewNoteUseCase(noteRepository, logger)

	return noteUseCase
}
