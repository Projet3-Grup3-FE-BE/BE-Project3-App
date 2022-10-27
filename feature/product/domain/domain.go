package domain

import (
	userDom "be_project3team3/feature/user/domain"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	Product_Name   string
	Description    string
	Price          int
	Image_Url      string
	Stock          int
	Shop_Name      string
	Category       string
	Id_User_Seller uint
}

type RepositoryInterface interface {
	GetAll() ([]Core, error)
	GetAllByCategory(category string) ([]Core, error)
	GetShop(id_user_seller string) ([]Core, error)
	GetShopByCategory(id_user_seller string, category string) ([]Core, error)
	Get(ID string) (Core, error)
	GetUser(idUser uint) (userDom.Core, error)
	// GetProductAllComment(ID string) (Core, []domComment.Core, error)
	Insert(newData Core) (Core, error)
	Update(updatedData Core, ID uint) (Core, error)
	Delete(ID string) (Core, error)
}
type ServiceInterface interface {
	GetAll(id_user_seller string, category string) ([]Core, error)
	Get(ID string) (Core, error)
	// GetProductAllComment(ID string) (Core, []domComment.Core, error)
	Insert(newData Core, c echo.Context) (Core, error)
	Update(updatedData Core, ID string, c echo.Context) (Core, error)
	Delete(ID string) (Core, error)
	IsAuthorized(c echo.Context) error
}
