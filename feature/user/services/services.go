package services

import (
	"errors"
	"strings"
	"time"

	"be_project3team3/config"
	"be_project3team3/feature/user/domain"
	"be_project3team3/helper"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

var key string

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func New(repo domain.Repository) domain.Service {
	return &repoService{
		qry: repo,
	}
}

type repoService struct {
	qry domain.Repository
}

// Register implements domain.Service
func (rs *repoService) Register(newUser domain.Core, c echo.Context) (domain.Core, error) {
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

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := helper.UploadProfile(c)
		if err != nil {
			return domain.Core{}, errors.New("Registration Failed. Cannot Upload Data.")
		}
		log.Print(res)
		newUser.ImageUrl = res
	} else {
		newUser.ImageUrl = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/dummy-profile-pic.png"
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

func (rs *repoService) LoginUser(newUser domain.Core) (domain.Core, error) {
	res, err := rs.qry.Login(newUser)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	// token := GenerateToken(res.ID)
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(newUser.Password))
	if err != nil {
		return domain.Core{}, errors.New("password tidak cocok")
	}
	return res, nil

}
func (rs *repoService) GenerateToken(id uint) string {
	claim := make(jwt.MapClaims)
	claim["authorized"] = true
	claim["id"] = id
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte(key))
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	return str
}

func (rs *repoService) UpdateProfile(updatedData domain.Core, c echo.Context) (domain.Core, error) {
	if updatedData.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(updatedData.Password), 10)
		updatedData.Password = string(generate)
	}

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := helper.UploadProfile(c)
		if err != nil {
			return domain.Core{}, errors.New("Registration Failed. Cannot Upload Data.")
		}
		log.Print(res)
		updatedData.ImageUrl = res
	} else {
		updatedData.ImageUrl = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/dummy-profile-pic.png"
	}

	res, err := rs.qry.Update(updatedData)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("rejected from database")
	}

	return res, nil
}

func (us *repoService) Delete(ID uint) error {
	err := us.qry.Delete(ID)
	if err != nil {
		log.Error(err.Error())
		return errors.New("no data")
	}

	return nil
}

func (rs *repoService) GetUser(getuserdata domain.Core) (domain.Core, error) {

	res, err := rs.qry.GetUser(getuserdata)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("rejected from database")
	}

	return res, nil
}

func (us *repoService) GetMe(ID uint) (domain.Core, error) {
	res, err := us.qry.GetMe(ID)
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("no data")
	}

	return res, nil
}
