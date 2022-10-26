package delivery

import (
	"be_project3team3/feature/user/domain"
)

func SuccessDelete(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}
func SuccessLogin(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

type registerRespons struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	ShopName string `json:"shop_name"`
	ImageUrl string `json:"image_url"`
}

type loginRespons struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	ShopName string `json:"shop_name"`
	ImageUrl string `json:"image_url"`
	Token    string `json:"token"`
}

type UpdateRespons struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	ShopName string `json:"shop_name"`
	ImageUrl string `json:"image_url"`
}

type GetDataRespons struct {
	ID                uint   `json:"id"`
	Username          string `json:"username"`
	Name              string `json:"name"`
	Phone             string `json:"phone"`
	Recipient_address string `json:"recipient_address"`
	ShopName          string `json:"shop_name"`
}

type GetMeRespons struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	ShopName string `json:"shop_name"`
	ImageUrl string `json:"image_url"`
}

func ToResponseLogin(core interface{}, token string, code string) interface{} {
	var res interface{}
	switch code {
	case "login":
		cnv := core.(domain.Core)
		res = loginRespons{
			ID:       cnv.ID,
			Username: cnv.Username,
			Email:    cnv.Email,
			Password: cnv.Password,
			Name:     cnv.Name,
			Phone:    cnv.Phone,
			Address:  cnv.Address,
			ShopName: cnv.ShopName,
			ImageUrl: cnv.ImageUrl,
			Token:    token,
		}
	}
	return res

}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = registerRespons{
			ID:       cnv.ID,
			Username: cnv.Username,
			Email:    cnv.Email,
			Password: cnv.Password,
			Name:     cnv.Name,
			Phone:    cnv.Phone,
			Address:  cnv.Address,
			ShopName: cnv.ShopName,
			ImageUrl: cnv.ImageUrl,
		}
	case "upd":
		cnv := core.(domain.Core)
		res = UpdateRespons{
			ID:       cnv.ID,
			Username: cnv.Username,
			Email:    cnv.Email,
			Password: cnv.Password,
			Name:     cnv.Name,
			Phone:    cnv.Phone,
			Address:  cnv.Address,
			ShopName: cnv.ShopName,
			ImageUrl: cnv.ImageUrl,
		}
	case "getuser":
		cnv := core.(domain.Core)
		res = GetDataRespons{
			ID:                cnv.ID,
			Username:          cnv.Username,
			Name:              cnv.Name,
			Phone:             cnv.Phone,
			Recipient_address: cnv.Recipient_address,
			ShopName:          cnv.ShopName,
		}
	case "getMe":
		cnv := core.(domain.Core)
		res = GetMeRespons{
			ID:       cnv.ID,
			Username: cnv.Username,
			Email:    cnv.Email,
			Password: cnv.Password,
			Name:     cnv.Name,
			Phone:    cnv.Phone,
			Address:  cnv.Address,
			ShopName: cnv.ShopName,
			ImageUrl: cnv.ImageUrl,
		}
	}
	return res
}
