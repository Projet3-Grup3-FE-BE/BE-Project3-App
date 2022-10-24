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

type loginRespons struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Bio   string `json:"bio"`
}

type UpdateRespons struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
	Location string `json:"location"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = registerRespons{
			ID:                cnv.ID,
			Username:          cnv.Username,
			Email:             cnv.Email,
			Password:          cnv.Password,
			Name:              cnv.Name,
			Alamat_pengiriman: cnv.Alamat_pengiriman,
		}
	case "login":
		cnv := core.(domain.Core)
		res = loginRespons{ID: cnv.ID, Email: cnv.Email, Name: cnv.Name, Phone: cnv.Phone, Bio: cnv.Bio}
	case "upd":
		cnv := core.(domain.Core)
		res = UpdateRespons{ID: cnv.ID, Email: cnv.Email, Name: cnv.Name, Phone: cnv.Phone, Bio: cnv.Bio, Gender: cnv.Gender, Location: cnv.Location}
	}
	return res
}
