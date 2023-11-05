package user

import (
	"github.com/aldhipradana/apidemo/schemas"
	"gorm.io/gorm"
)

type Usecases struct {
	Db *gorm.DB
}

func (u *Usecases) Create(user *schemas.User) error {
	return u.Db.Create(user).Error
}

func (u *Usecases) GetById(id string) (*schemas.User, error) {
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
	return u.Db.Save(user).Error
}

func (u *Usecases) Delete(id uint) error {
	return u.Db.Delete(&schemas.User{}, id).Error
}
