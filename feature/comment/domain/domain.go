package domain

import (
	userDom "be_project3team3/feature/user/domain"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	Name_User     string
	Comment_Value string
	IDUser        uint
	IDPosting     uint
}

type RepositoryInterface interface {
	Get(ID string) (Core, error)
	GetAll(idPosting string) ([]Core, error)
	GetUser(idUser uint) (userDom.Core, error)
	Insert(newData Core) (Core, error)
	Update(updatedData Core, ID uint) (Core, error)
	Delete(ID string) (Core, error)
}
type ServiceInterface interface {
	GetAll(idPosting string) ([]Core, error)
	Get(ID string) (Core, error)
	Insert(newData Core, c echo.Context) (Core, error)
	Update(updatedData Core, ID string, c echo.Context) (Core, error)
	Delete(ID string) (Core, error)
	IsAuthorized(c echo.Context) error
}
