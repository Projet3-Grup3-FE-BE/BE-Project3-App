package services

import (
	"be_project3team3/feature/comment/domain"
	"be_project3team3/utils/jwt"
	"errors"
	loggo "log"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type commentService struct {
	qry domain.RepositoryInterface
}

func New(repo domain.RepositoryInterface) domain.ServiceInterface {
	return &commentService{
		qry: repo,
	}
}

func (bs *commentService) Insert(newData domain.Core, c echo.Context) (domain.Core, error) {
	if IsEmptyValidation(newData) {
		return domain.Core{}, errors.New("Failed. New data empty. ")
	}
	idUser := jwt.ExtractIdToken(c)
	userData, err := bs.qry.GetUser(idUser)
	if err != nil {
		return domain.Core{}, errors.New("Failed. User not found. ")
	}

	newData.Name_User = userData.Name
	newData.IDUser = idUser

	res, err := bs.qry.Insert(newData)

	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("Failed. Duplicate record. Rejected from database")
		} else {
			return domain.Core{}, errors.New("Failed. Some problem on database")
		}
	}

	return res, nil
}

func (bs *commentService) Update(updatedData domain.Core, idComment string, c echo.Context) (domain.Core, error) {
	loggo.Println("\n\n\nisi id", updatedData.ID)
	idUint, _ := strconv.Atoi(idComment)
	if idUint == 0 {
		return domain.Core{}, errors.New("Failed. Data id is empty.")
	} else {
		// validasi data yang di update ada atau ngga by id
		loggo.Println("\n\n\n\nValidate update data success, id update", updatedData.ID)

		// loggo.Println("\n\n\n\n error1", err.Error(), "\n\n\nerror2 :", err2.Error())
		if _, err := bs.qry.Get(idComment); err != nil {
			return domain.Core{}, errors.New("Failed. Data not found. Add first.")
		} else {
			idUser := jwt.ExtractIdToken(c)
			userData, err := bs.qry.GetUser(idUser)
			if err != nil {
				return domain.Core{}, errors.New("Failed. User not found. ")
			}

			updatedData.Name_User = userData.Name
			updatedData.IDUser = idUser

			resUpdate, err := bs.qry.Update(updatedData, uint(idUint))

			if err != nil {
				log.Error(err.Error())
				if strings.Contains(err.Error(), "column") {
					return domain.Core{}, errors.New("Failed. Rejected from database")
				} else if strings.Contains(err.Error(), "found") {
					return domain.Core{}, errors.New("Failed. Fata record not found. add first.")
				} else {
					return domain.Core{}, errors.New("Failed. Some problem on database")
				}
			}
			return resUpdate, err
		}
	}
}

func IsEmptyValidation(data domain.Core) bool {
	result := false
	if data == (domain.Core{}) {
		result = true
	}
	return result
}
func UintToString(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

func GetValidation(res domain.Core, err error) (domain.Core, error) {
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("Failed. Database error.")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("Failed. No data.")
		} else {
			return domain.Core{}, errors.New("Failed. Some problem on database.")
		}
	}

	return res, err
}

func (bs *commentService) Get(ID string) (domain.Core, error) {
	res, err := bs.qry.Get(ID)

	res2, err2 := GetValidation(res, err)

	return res2, err2
}

func (bs *commentService) GetAll(idPosting string) ([]domain.Core, error) {
	res, err := bs.qry.GetAll(idPosting)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("Failed. Database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("Failed. No data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("No Data")
	}

	return res, nil
}

func (bs *commentService) Delete(idPosting string) (domain.Core, error) {
	res, err := bs.qry.Delete(idPosting)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		} else {
			return domain.Core{}, errors.New("some problem on database")
		}
	}

	return res, nil
}

func (bs *commentService) IsAuthorized(c echo.Context) error {
	id, exp := jwt.ExtractToken(c)

	if id == 0 {
		return errors.New("Request not authorized. Please check token. User not found.")
	} else if time.Now().Unix() > exp {
		return errors.New("Request not authorized. Please check token. Expired token.")
	} else {
		return nil
	}
}
