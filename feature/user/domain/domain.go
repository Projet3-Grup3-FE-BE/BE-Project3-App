package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID                uint
	Username          string
	Email             string
	Password          string
	Name              string
	Alamat_pengiriman string
	Phone             string
	Bio               string
	Gender            string
	Location          string
}

type Repository interface {
	Insert(newUser Core) (Core, error) //register
	Login(newUser Core) (Core, error)
	Delete(ID uint) error
	Update(updateData Core) (Core, error)
}

type Service interface {
	Register(newUser Core) (Core, error)
	LoginUser(newUser Core) (Core, error)
	GenerateToken(id uint) string
	Delete(id uint) error
	UpdateProfile(updateData Core) (Core, error)
}
type Handler interface {
	Register() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
	DeleteByID() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
}
