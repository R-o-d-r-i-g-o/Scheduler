package model

import (
	db "scheduler/database"
	"scheduler/module"
)

type IUserReferences interface {
	Find() error
	Create() error
	Update() error
	Delete() error
}

type UserReferences struct {
	IUserReferences
	User  *module.User
	Users *[]module.User
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
	err = db.GetGormDB().Where("cpf_cnpj = ?", u.User).Updates(&u.User).Error
	return
}

func (u *UserReferences) Delete() (err error) {
	err = db.GetGormDB().Where("cpf_cnpj = ?", u.User).Delete(&u.User).Error
	return
}
