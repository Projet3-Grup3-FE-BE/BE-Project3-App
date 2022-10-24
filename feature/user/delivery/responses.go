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
func SuccessLogin(msg string, token string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
		"token":   token,
	}
}

type registerRespons struct {
	ID                uint   `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Name              string `json:"name"`
	Alamat_pengiriman string `json:"alamat_pengiriman"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = registerRespons{
			Username:          cnv.Username,
			Email:             cnv.Email,
			Password:          cnv.Password,
			Name:              cnv.Name,
			Alamat_pengiriman: cnv.Alamat_pengiriman,
		}
	}
	return res
}
