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

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Username: cnv.Username, Email: cnv.Email, Password: cnv.Password, Name: cnv.Name, Alamat_pengiriman: cnv.Alamat_pengiriman}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}

	}
	return domain.Core{}
}
