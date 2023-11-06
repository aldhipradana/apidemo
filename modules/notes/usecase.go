package notes

import (
	"github.com/aldhipradana/apidemo/schemas"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Usecases struct {
	Db *gorm.DB
}

func (u *Usecases) Create(note *schemas.Notes) error {

	note.ID = uuid.NewV4()

	return u.Db.Create(note).Error
}

func (u *Usecases) GetById(id uuid.UUID) (*schemas.Notes, error) {
	var note schemas.Notes
	err := u.Db.First(&note, id).Error
	return &note, err
}

func (u *Usecases) List() ([]schemas.Notes, error) {
	var notes []schemas.Notes
	err := u.Db.Find(&notes).Error
	return notes, err
}

func (u *Usecases) Update(note *schemas.Notes) error {
	oldNote, err := u.GetById(note.ID)
	if err != nil {
		return err
	}

	note.CreatedAt = oldNote.CreatedAt
	note.UpdatedAt = oldNote.UpdatedAt

	// dont update nil value
	return u.Db.Model(note).Updates(note).Error
}

func (u *Usecases) Delete(id uuid.UUID) error {
	return u.Db.Delete(&schemas.Notes{}, id).Error
}
