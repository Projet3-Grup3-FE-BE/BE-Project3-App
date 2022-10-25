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
	Phone             string
	Address           string
	ShopName          string
	ImageUrl          string
	Recipient_address string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:             gorm.Model{ID: du.ID},
		Username:          du.Username,
		Email:             du.Email,
		Password:          du.Password,
		Name:              du.Name,
		Phone:             du.Phone,
		Address:           du.Address,
		ShopName:          du.ShopName,
		ImageUrl:          du.ImageUrl,
		Recipient_address: du.Recipient_address,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:                u.ID,
		Username:          u.Username,
		Email:             u.Email,
		Password:          u.Password,
		Name:              u.Name,
		Phone:             u.Phone,
		Address:           u.Address,
		ShopName:          u.ShopName,
		ImageUrl:          u.ImageUrl,
		Recipient_address: u.Recipient_address,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Username: val.Username, Email: val.Email, Password: val.Password, Name: val.Name, Phone: val.Phone, Address: val.Address, ShopName: val.ShopName, ImageUrl: val.ImageUrl, Recipient_address: val.Recipient_address})
	}
	return res
}
