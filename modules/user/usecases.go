package user

import (
	"github.com/aldhipradana/apidemo/schemas"
	"github.com/aldhipradana/apidemo/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Usecases struct {
	Db *gorm.DB
}

func (u *Usecases) Create(user *schemas.User) error {

	user.ID = uuid.NewV4()

	user.Password = utils.HashMD5(user.Password)

	return u.Db.Create(user).Error
}

func (u *Usecases) GetById(id uuid.UUID) (*schemas.User, error) {
	var user schemas.User
	err := u.Db.First(&user, id).Error
	return &user, err
}

func (u *Usecases) List() ([]schemas.User, error) {
	var users []schemas.User
	err := u.Db.Find(&users).Error
	return users, err
}

func (u *Usecases) Update(user *schemas.User) error {
	us, err := u.GetById(user.ID)
	if err != nil {
		return err
	}

	user.CreatedAt = us.CreatedAt
	user.UpdatedAt = us.UpdatedAt
	user.Password = utils.HashMD5(user.Password)

	// dont update nil value
	return u.Db.Model(user).Updates(user).Error
}

func (u *Usecases) Delete(id uuid.UUID) error {
	return u.Db.Delete(&schemas.User{}, id).Error
}
