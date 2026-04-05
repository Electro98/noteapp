package note

import (
	"errors"
	"net/http"

	"github.com/electro98/noteapp/domain"
	"github.com/electro98/noteapp/utils"
	"github.com/labstack/echo/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type noteHandler struct {
	noteUseCase domain.NoteUseCase
	logger      *zerolog.Logger
}

func NewNoteHandler(g *echo.Group, ctx domain.NoteUseCase, logger *zerolog.Logger) {
	h := &noteHandler{noteUseCase: ctx, logger: logger}

	g.GET("/note", h.AllNotes)
	g.GET("/note/:id", h.GetNote)
	g.POST("/note", h.CreateNote)
	g.PUT("/note", h.UpdateNote)
	g.DELETE("/note", h.DeleteNote)
}

func (h noteHandler) AllNotes(c *echo.Context) error {
	var filter domain.NoteFilterParam
	if err := c.Bind(&filter); err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/AllNotes/BindQueryParams")
		return c.JSON(http.StatusBadRequest, utils.JsonMessage("Bad Request"))
	}
	notes, err := h.noteUseCase.List(c.Request().Context(), filter)
	if err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/AllNotes/List")
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, notes)
}

func (h noteHandler) CreateNote(c *echo.Context) error {
	var newNote domain.NoteNewParam
	if err := c.Bind(&newNote); err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/CreateNote/BindBody")
		// return c.JSON(http.StatusBadRequest, utils.JsonMessage("Bad Request"))
		return err // Wrap it??
	}
	note, err := h.noteUseCase.Create(c.Request().Context(), newNote)
	if err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/CreateNote/Create")
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, note)
}

func (h noteHandler) UpdateNote(c *echo.Context) error {
	var updateNote domain.NoteUpdateParam
	if err := c.Bind(&updateNote); err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/UpdateNote/BindBody")
		return err
	}
	note, err := h.noteUseCase.Update(c.Request().Context(), updateNote)
	if err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/UpdateNote/Update")
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, note)
}

func (h noteHandler) GetNote(c *echo.Context) error {
	id, err := echo.PathParam[uint](c, "id")
	if err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/GetNote/PathParam")
		return c.JSON(http.StatusBadRequest, utils.JsonMessage("Bad Request"))
	}
	note, err := h.noteUseCase.GetNoteById(c.Request().Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusForbidden)
	} else if err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/GetNote/GetNoteById")
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, note)
}

func (h noteHandler) DeleteNote(c *echo.Context) error {
	var noteDelete domain.NoteDeleteParam
	if err := c.Bind(&noteDelete); err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/DeleteNote/Bind")
		return c.JSON(http.StatusBadRequest, utils.JsonMessage("Bad Request"))
	}
	err := h.noteUseCase.Delete(c.Request().Context(), noteDelete)
	if errors.Is(err, utils.ErrNoFoundInDB) {
		return c.NoContent(http.StatusForbidden)
	} else if err != nil {
		h.logger.Error().Err(err).Msg("noteHandler/DeleteNote/Delete")
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
