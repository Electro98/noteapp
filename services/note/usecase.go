package note

import (
	"context"
	"fmt"

	"github.com/electro98/noteapp/domain"
	"github.com/electro98/noteapp/models"
	"github.com/electro98/noteapp/utils"
	"github.com/rs/zerolog"
)

type noteUseCase struct {
	noteRepository domain.NoteRepository
	logger         *zerolog.Logger
}

func (n noteUseCase) List(ctx context.Context, param domain.NoteFilterParam) ([]models.Note, error) {
	param.Validate() // TODO: Think about it. I don't like void functions
	notes, err := n.noteRepository.List(ctx, param)
	if err != nil {
		n.logger.Error().Err(err).Msg("noteUseCase/List/List")
		return nil, fmt.Errorf("noteUseCase/List: %w", err)
	}
	return notes, err
}

func (n noteUseCase) GetNoteById(ctx context.Context, id uint) (models.Note, error) {
	note, err := n.noteRepository.GetNoteById(ctx, id)
	if err != nil {
		n.logger.Error().Err(err).Msg("noteUseCase/GetNoteById/GetNoteById")
		return models.Note{}, fmt.Errorf("noteUseCase/GetNoteById: %w", err)
	}
	return note, nil
}

func (n noteUseCase) Create(ctx context.Context, param domain.NoteNewParam) (models.Note, error) {
	note := &models.Note{
		Title:   param.Title,
		Content: param.Content,
	}
	note, err := n.noteRepository.Create(ctx, note)
	if err != nil {
		n.logger.Error().Err(err).Msg("noteUseCase/Create/Create")
		return models.Note{}, fmt.Errorf("noteUseCase/Create: %w", err)
	}
	return *note, nil
}

func (n noteUseCase) Update(ctx context.Context, param domain.NoteUpdateParam) (models.Note, error) {
	note := models.Note{
		BaseModel: utils.BaseModel{ID: param.ID},
		Title:     param.Title,
		Content:   param.Content,
	}
	err := n.noteRepository.Update(ctx, &note)
	if err != nil {
		n.logger.Error().Err(err).Msg("noteUseCase/Update/Update")
		return models.Note{}, fmt.Errorf("noteUseCase/Update: %w", err)
	}
	note, err = n.noteRepository.GetNoteById(ctx, param.ID)
	if err != nil {
		n.logger.Error().Err(err).Msg("noteUseCase/Update/GetNoteById")
		return models.Note{}, fmt.Errorf("noteUseCase/Update: %w", err)
	}
	return note, nil
}

func (n noteUseCase) Delete(ctx context.Context, param domain.NoteDeleteParam) error {
	err := n.noteRepository.DeleteNoteById(ctx, param.ID)
	if err != nil {
		n.logger.Error().Err(err).Msg("noteUseCase/Delete/DeleteNoteById")
		return fmt.Errorf("noteUseCase/Delete: %w", err)
	}
	return nil
}

func NewNoteUseCase(noteRep domain.NoteRepository, logger *zerolog.Logger) domain.NoteUseCase {
	return &noteUseCase{
		noteRepository: noteRep,
		logger:         logger,
	}
}
