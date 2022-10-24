package services

import (
	"be_project3team3/config"
	"be_project3team3/feature/user/domain"
	"be_project3team3/utils/jwt"
	"errors"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}

	newUser.Password = string(generate)
	res, err := rs.qry.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New(config.DUPLICATED_DATA)
		}

		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil

}

// Login implements domain.Service
func (rs *repoService) Login(email, password string) (domain.Core, string, error) {

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return domain.Core{}, "", errors.New("Email or password empty")
	}

	res, err := rs.qry.GetUser(email)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, "", errors.New("Failed. Error database.")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, "", errors.New("Failed. Email or Password not found.")
		} else {
			return domain.Core{}, "", errors.New("Failed. Process error. Please contact Admin")
		}
	} else {
		// loggo.Println("res pass", res.Password, "\n\npass", password)
		err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
		if err != nil {
			return domain.Core{}, "", errors.New("Failed. Incorrect Password.")
		}

		token, err := jwt.GenerateJWTToken(res.ID)
		if err != nil {
			return domain.Core{}, "", err
		}
		return res, token, nil
	}

}

// DeleteProfile implements domain.Service
// func (*repoService) DeleteProfile(ID uint) (domain.Core, error) {
// 	panic("unimplemented")
// }

// UpdateProfile implements domain.Service
func (rs *repoService) UpdateProfile(updatedData domain.Core, c echo.Context) (domain.Core, error) {
	userId, _ := jwt.ExtractToken(c)
	log.Printf("\n\n\nisi service = ", updatedData, "\n\n\n")

	generate, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}
	updatedData.Password = string(generate)

	res, err := rs.qry.Update(updatedData, userId)
	log.Printf("\n\n\nisi service 2 = ", res, "\n\n\n")
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.Core{}, gorm.ErrRecordNotFound
		} else {
			return domain.Core{}, errors.New(config.DATABASE_ERROR)
		}
	}
	return res, nil
}

func (rs *repoService) IsAuthorized(c echo.Context) error {
	id, exp := jwt.ExtractToken(c)
	// loggo.Println("id dr tken = ", id)
	// loggo.Println("exp dr tken = ", exp)
	if id == 0 {
		return errors.New("Request not authorized. Please check token. User not found.")
	} else if time.Now().Unix() > exp {
		return errors.New("Request not authorized. Please check token. Expired token.")
	} else {
		return nil
	}
}

func (rs *repoService) DeleteProfile(c echo.Context) (domain.Core, error) {
	id := jwt.ExtractIdToken(c)
	res, err := rs.qry.Delete(id)
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return res, gorm.ErrRecordNotFound
		} else {
			return res, errors.New(config.DATABASE_ERROR)
		}
	}
	return domain.Core{}, nil
}

// ShowAllUser implements domain.Service
func (rs *repoService) ShowAllUser() ([]domain.Core, error) {
	res, err := rs.qry.GetAll()

	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New(config.DATA_NOTFOUND)
	}
	return res, nil
}

// Profile implements domain.Service
func (rs *repoService) Profile(Email string) (domain.Core, error) {
	res, err := rs.qry.Get(Email)
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.Core{}, gorm.ErrRecordNotFound
		} else {
			return domain.Core{}, errors.New(config.DATABASE_ERROR)
		}
	}
	return res, nil
}

// // check update password
// func  (rs *repoService) isPasswordMatch(email string, newPassword string) (domain.Core, error) {
// 	if strings.TrimSpace(email) == "" || strings.TrimSpace(newPassword) == "" {
// 		return domain.Core{}, errors.New("Email or password empty")
// 	}

// 	res, err := rs.qry.GetUser(email)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if strings.Contains(err.Error(), "table") {
// 			return domain.Core{}, errors.New("Failed. Error database.")
// 		} else if strings.Contains(err.Error(), "found") {
// 			return domain.Core{}, errors.New("Failed. Email or Password not found.")
// 		} else {
// 			return domain.Core{}, errors.New("Failed. Process error. Please contact Admin")
// 		}
// 	} else {
// 		// loggo.Println("res pass", res.Password, "\n\npass", password)
// 		err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(newPassword))
// 		if err != nil {
// 			return domain.Core{},errors.New("Failed. Incorrect Password.")
// 		}
// 	}
// 	return res, nil
// }
