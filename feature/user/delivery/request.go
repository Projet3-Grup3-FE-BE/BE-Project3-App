package delivery

import (
	"be_project3team3/feature/user/domain"
)

type RegisterFormat struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	ShopName string `json:"shop_name" form:"shop_name"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

type UpdateFormat struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	ShopName string `json:"shop_name" form:"shop_name"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

type GetUserFormat struct {
	ID                uint   `json:"id" form:"id"`
	Username          string `json:"username" form:"username"`
	Name              string `json:"name" form:"name"`
	Phone             string `json:"phone" form:"phone"`
	Recipient_address string `json:"recipient_address" form:"recipient_address"`
	ShopName          string `json:"shop_name" form:"shop_name"`
}

type GetMeFormat struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	ShopName string `json:"shop_name" form:"shop_name"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Username: cnv.Username, Email: cnv.Email, Password: cnv.Password, Name: cnv.Name, Phone: cnv.Phone, Address: cnv.Address, ShopName: cnv.ShopName, ImageUrl: cnv.ImageUrl}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Password: cnv.Password, Name: cnv.Name, Phone: cnv.Phone, Address: cnv.Address, ShopName: cnv.ShopName, ImageUrl: cnv.ImageUrl}
	case GetUserFormat:
		cnv := i.(GetUserFormat)
		return domain.Core{ID: cnv.ID, Username: cnv.Username, Name: cnv.Name, Phone: cnv.Phone, Recipient_address: cnv.Recipient_address, ShopName: cnv.ShopName}
	case GetMeFormat:
		cnv := i.(GetMeFormat)
		return domain.Core{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Password: cnv.Password, Name: cnv.Name, Phone: cnv.Phone, Address: cnv.Address, ShopName: cnv.ShopName, ImageUrl: cnv.ImageUrl}
	}
	return domain.Core{}
}
