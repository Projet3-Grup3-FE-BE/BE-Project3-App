package repository

import (
	"be_project3team3/feature/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string
	Email             string
	Password          string
	Name              string
	Alamat_pengiriman string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:             gorm.Model{ID: du.ID},
		Username:          du.Username,
		Email:             du.Email,
		Password:          du.Password,
		Name:              du.Name,
		Alamat_pengiriman: du.Alamat_pengiriman,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:                u.ID,
		Username:          u.Username,
		Email:             u.Email,
		Password:          u.Password,
		Name:              u.Name,
		Alamat_pengiriman: u.Alamat_pengiriman,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Username: val.Username, Email: val.Email, Password: val.Password, Name: val.Name, Alamat_pengiriman: val.Alamat_pengiriman})
	}
	return res
}
