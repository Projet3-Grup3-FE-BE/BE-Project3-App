package delivery

import "be_project3team3/feature/user/domain"

type RegisterFormat struct {
	Username          string `json:"username" form:"username"`
	Email             string `json:"email" form:"email"`
	Password          string `json:"password"  form:"password"`
	Name              string `json:"name" form:"name"`
	Alamat_pengiriman string `json:"alamat_pengiriman" form:"alamat_pengiriman"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

type UpdateFormat struct {
	ID       uint   `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Bio      string `json:"bio" form:"bio"`
	Gender   string `json:"gender" form:"gender"`
	Location string `json:"location" form:"location"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Username: cnv.Username, Email: cnv.Email, Password: cnv.Password, Name: cnv.Name, Alamat_pengiriman: cnv.Alamat_pengiriman}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Email: cnv.Email, Password: cnv.Password, Name: cnv.Name, Phone: cnv.Phone, Bio: cnv.Bio, Gender: cnv.Gender, Location: cnv.Location}
	}
	return domain.Core{}
}
