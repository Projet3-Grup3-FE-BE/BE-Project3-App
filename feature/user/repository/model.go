package repository

import (
	"be_project3team3/feature/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Bio      string
	Gender   string
	Location string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Name:     du.Name,
		Email:    du.Email,
		Password: du.Password,
		Phone:    du.Phone,
		Bio:      du.Bio,
		Gender:   du.Gender,
		Location: du.Location,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Bio:      u.Bio,
		Gender:   u.Gender,
		Location: u.Location,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{
			ID:       val.ID,
			Name:     val.Name,
			Email:    val.Email,
			Password: val.Password,
			Phone:    val.Phone,
			Bio:      val.Bio,
			Gender:   val.Gender,
			Location: val.Location,
		})
	}
	return res
}
