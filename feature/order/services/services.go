package services

import (
	"be_project3team3/feature/order/domain"
	"be_project3team3/utils/jwt"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type orderService struct {
	qry domain.RepositoryInterface
}

func New(repo domain.RepositoryInterface) domain.ServiceInterface {
	return &orderService{
		qry: repo,
	}
}

func (bs *orderService) Insert(newData domain.Core, c echo.Context) (domain.Core, error) {
	if IsEmptyValidation(newData) {
		return domain.Core{}, errors.New("Failed. New data empty. ")
	}
	idUser := jwt.ExtractIdToken(c)
	newData.IdUserBuyer = idUser
	newData.OrderStatus = "process_order"

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

// func (bs *orderService) Update(updatedData domain.Core, idorder string, c echo.Context) (domain.Core, error) {
// 	loggo.Println("\n\n\nisi id", updatedData.ID)
// 	idorderUint, _ := strconv.Atoi(idorder)
// 	if idorderUint == 0 {
// 		return domain.Core{}, errors.New("Failed. Data id is empty.")
// 	} else {
// 		// validasi data yang di update ada atau ngga by id
// 		loggo.Println("\n\n\n\nValidate update data success, id update", updatedData.ID)

// 		// loggo.Println("\n\n\n\n error1", err.Error(), "\n\n\nerror2 :", err2.Error())
// 		if _, err := bs.qry.Get(idorder); err != nil {
// 			return domain.Core{}, errors.New("Failed. Data not found. Add first.")
// 		} else {
// 			idUser := jwt.ExtractIdToken(c)
// 			userData, err := bs.qry.GetUser(idUser)
// 			if err != nil {
// 				return domain.Core{}, errors.New("Failed. User not found. ")
// 			}

// 			updatedData.Shop_Name = userData.ShopName
// 			updatedData.Id_User_Seller = idUser
// 			resUpdate, err := bs.qry.Update(updatedData, uint(idorderUint))

// 			if err != nil {
// 				log.Error(err.Error())
// 				if strings.Contains(err.Error(), "column") {
// 					return domain.Core{}, errors.New("Failed. Rejected from database")
// 				} else if strings.Contains(err.Error(), "found") {
// 					return domain.Core{}, errors.New("Failed. Fata record not found. add first.")
// 				} else {
// 					return domain.Core{}, errors.New("Failed. Some problem on database")
// 				}
// 			}
// 			return resUpdate, err
// 		}
// 	}
// }

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

func (bs *orderService) Get(ID string) (domain.Core, error) {
	res, err := bs.qry.Get(ID)

	res2, err2 := GetValidation(res, err)

	return res2, err2
}

// func (bs *orderService) GetAll(id_user_seller string, category string) ([]domain.Core, error) {
// 	var res []domain.Core
// 	var err error
// 	if id_user_seller == "" && category == "" {
// 		res, err = bs.qry.GetAll()
// 	} else if id_user_seller == "" && category != "" {
// 		res, err = bs.qry.GetAllByCategory(category)

// 	} else if id_user_seller != "" && category == "" {
// 		res, err = bs.qry.GetShop(id_user_seller)

// 	} else if id_user_seller != "" && category != "" {
// 		res, err = bs.qry.GetShopByCategory(id_user_seller, category)
// 	}

// 	if err != nil {
// 		log.Error(err.Error())
// 		if strings.Contains(err.Error(), "table") {
// 			return nil, errors.New("Failed. Database error")
// 		} else if strings.Contains(err.Error(), "found") {
// 			return nil, errors.New("Failed. No data")
// 		}
// 	}

// 	if len(res) == 0 {
// 		log.Info("no data")
// 		return nil, errors.New("no data")
// 	}

// 	return res, nil
// }

// func (bs *productService) GetproductAllComment(ID string) (domain.Core, []domComment.Core, error) {
// 	resproduct, resComments, err := bs.qry.GetproductAllComment(ID)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if strings.Contains(err.Error(), "table") {
// 			return domain.Core{}, nil, errors.New("Failed. Database error")
// 		} else if strings.Contains(err.Error(), "found") {
// 			return domain.Core{}, nil, errors.New("Failed. No data")
// 		} else {
// 			log.Print("Log :", err.Error())
// 			loggo.Println("Log : ", err.Error())
// 			return domain.Core{}, nil, errors.New("Failed. Please check log.")
// 		}
// 	}

// 	return resproduct, resComments, nil
// }

func (bs *orderService) Delete(idorder string) (domain.Core, error) {
	res, err := bs.qry.Delete(idorder)
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

func (bs *orderService) IsAuthorized(c echo.Context) error {
	id, exp := jwt.ExtractToken(c)

	if id == 0 {
		return errors.New("Request not authorized. Please check token. User not found.")
	} else if time.Now().Unix() > exp {
		return errors.New("Request not authorized. Please check token. Expired token.")
	} else {
		return nil
	}
}
