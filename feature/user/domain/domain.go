package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Phone    string
	Bio      string
	Gender   string
	Location string
}

type Repository interface {
	Insert(newUser Core) (Core, error)  //register
	GetUser(email string) (Core, error) //login
	Update(updatedData Core, ID uint) (Core, error)
	Get(email string) (Core, error)
	Delete(ID uint) (Core, error)
	GetAll() ([]Core, error)
}

type Service interface {
	Register(newUser Core) (Core, error)
	Login(email, password string) (Core, string, error)
	UpdateProfile(updatedData Core, c echo.Context) (Core, error)
	Profile(email string) (Core, error)
	DeleteProfile(c echo.Context) (Core, error)
	IsAuthorized(c echo.Context) error
	ShowAllUser() ([]Core, error)
}
