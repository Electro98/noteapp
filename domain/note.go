package domain

import (
	"context"

	"github.com/electro98/noteapp/models"
	"github.com/electro98/noteapp/utils"
)

type (
	NoteNewParam struct {
		Title   string `json:"title" xml:"title" form:"title"`
		Content string `json:"content" xml:"content" form:"content"`
	}

	NoteUpdateParam struct {
		ID      uint   `json:"id" xml:"id" form:"id"`
		Title   string `json:"title" xml:"title" form:"title"`
		Content string `json:"content" xml:"content" form:"content"`
	}

	NoteDeleteParam struct {
		ID uint `query:"id" param:"id" json:"id" xml:"id" form:"id"`
	}

	NoteFilterParam struct {
		utils.Filter
	}

	NoteUseCase interface {
		List(ctx context.Context, param NoteFilterParam) ([]models.Note, error)
		GetNoteById(ctx context.Context, id uint) (models.Note, error)
		Create(ctx context.Context, param NoteNewParam) (models.Note, error)
		Update(ctx context.Context, param NoteUpdateParam) (models.Note, error)
		Delete(ctx context.Context, param NoteDeleteParam) error
	}

	NoteRepository interface {
		List(ctx context.Context, param NoteFilterParam) ([]models.Note, error)
		GetNoteById(ctx context.Context, id uint) (models.Note, error)
		Create(ctx context.Context, note *models.Note) (*models.Note, error)
		Update(ctx context.Context, note *models.Note) error
		DeleteNoteById(ctx context.Context, id uint) error
	}
)
