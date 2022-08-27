package user

import (
	db "scheduler/adapters/database"
	"scheduler/adapters/database/model"
)

type IUserReferences interface {
	Find() error
	Create() error
	Update() error
	Delete() error
}

type UserReferences struct {
	IUserReferences
	User  *model.User
	Users *[]model.User
}

func (u *UserReferences) Find() (err error) {
	err = db.GetGormDB().Find(&u.Users).Error
	return
}

func (u UserReferences) Create() (err error) {
	err = db.GetGormDB().Create(&u.User).Error
	return
}

func (u *UserReferences) Update() (err error) {
	err = db.GetGormDB().Where("id = ?", u.User).Updates(&u.User).Error
	return
}

func (u *UserReferences) Delete() (err error) {
	err = db.GetGormDB().Where("id = ?", u.User).Delete(&u.User).Error
	return
}
