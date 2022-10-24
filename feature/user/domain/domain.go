package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID                uint
	Username          string
	Email             string
	Password          string
	Name              string
	Alamat_pengiriman string
}

type Repository interface {
	Insert(newUser Core) (Core, error) //register

}

type Service interface {
	Register(newUser Core) (Core, error)
}
type Handler interface {
	Register() echo.HandlerFunc
}
