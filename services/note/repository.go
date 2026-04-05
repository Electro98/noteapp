package note

import (
	"context"

	"github.com/electro98/noteapp/domain"
	"github.com/electro98/noteapp/gen"
	"github.com/electro98/noteapp/models"
	"github.com/electro98/noteapp/utils"
	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

func (n noteRepository) List(ctx context.Context, param domain.NoteFilterParam) ([]models.Note, error) {
	notes, err := gorm.G[models.Note](n.db).
		Limit(int(param.Limit)).
		Offset(int(param.Offset)).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n noteRepository) GetNoteById(ctx context.Context, id uint) (models.Note, error) {
	note, err := gen.Query[models.Note](n.db).GetById(ctx, id)
	return note, err
}

func (n noteRepository) Create(ctx context.Context, note *models.Note) (*models.Note, error) {
	err := gorm.G[models.Note](n.db).Create(ctx, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n noteRepository) Update(ctx context.Context, note *models.Note) error {
	rows, err := gorm.G[models.Note](n.db).
		Where(gen.Note.ID.Eq(note.ID)).
		Updates(ctx, models.Note{
			Title:   note.Title,
			Content: note.Content,
		})
	if err != nil {
		return err
	} else if rows == 0 {
		return utils.ErrNoFoundInDB
	}
	return nil
}

func (n noteRepository) DeleteNoteById(ctx context.Context, id uint) error {
	rows, err := gorm.G[models.Note](n.db).Where(gen.Note.ID.Eq(id)).Delete(ctx)
	if rows == 0 && err == nil {
		return utils.ErrNoFoundInDB
	}
	return err
}

func NewNoteRepository(db *gorm.DB) domain.NoteRepository {
	return noteRepository{db: db}
}
