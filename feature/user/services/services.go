package services

import (
	"be_project3team3/config"
	"be_project3team3/feature/user/domain"

	// "be_project3team3/utils/jwt"
	"errors"

	// "time"

	// "github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
)

func New(repo domain.Repository) domain.Service {
	return &repoService{
		qry: repo,
	}
}

type repoService struct {
	qry domain.Repository
}

// Register implements domain.Service
func (rs *repoService) Register(newUser domain.Core) (domain.Core, error) {
	var temp string
	if newUser.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
		newUser.Password = string(generate)
	}
	if newUser.Password == "" {
		temp = newUser.Password
	} else {
		temp = "error"
	}
	res, err := rs.qry.Insert(newUser)

	if err != nil {
		if temp == "" {
			return domain.Core{}, errors.New("cannot encript password")
		}
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil

}
