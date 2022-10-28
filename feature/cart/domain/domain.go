package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID             uint
	Id_user        uint
	Id_product     uint
	Product_name   string
	Price          int
	ImageUrl       string
	Qty            int
	Sub_total      int
	Notes          string
	ShopName       string
	Id_user_seller uint
}

type Repository interface {
	Insert(newCart Core) (Core, error) //register
	GetCart(id uint) ([]Core, error)
	// GetAll() ([]Core, error)
	Delete(id uint) (Core, error)
	Update(updateCart Core) (Core, error)
}

type Service interface {
	AddCart(newCart Core) (Core, error)
	Delete(id uint) (Core, error)
	GetCart(id uint) ([]Core, error)
	UpdateCart(updateCartDetail Core) (Core, error)
	// ShowAllCart() ([]Core, error)
}
type Handler interface {
	AddDataCart() echo.HandlerFunc
	DeleteByID() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	GetCart() echo.HandlerFunc
	// ShowAllCart() echo.HandlerFunc
}
