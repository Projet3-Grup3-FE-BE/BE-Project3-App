package domain

import (
	domComment "be_project3team3/feature/comment/domain"
	userDom "be_project3team3/feature/user/domain"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Name_User string
	Image_Url string
	Content   string
	IDUser    uint
}

type RepositoryInterface interface {
	GetAll() ([]Core, error)
	Get(ID string) (Core, error)
	GetUser(idUser uint) (userDom.Core, error)
	GetPostingAllComment(ID string) (Core, []domComment.Core, error)
	Insert(newData Core) (Core, error)
	Update(updatedData Core, ID uint) (Core, error)
	Delete(ID string) (Core, error)
}
type ServiceInterface interface {
	GetAll() ([]Core, error)
	Get(ID string) (Core, error)
	GetPostingAllComment(ID string) (Core, []domComment.Core, error)
	Insert(newData Core, c echo.Context) (Core, error)
	Update(updatedData Core, ID string, c echo.Context) (Core, error)
	Delete(ID string) (Core, error)
	IsAuthorized(c echo.Context) error
}
