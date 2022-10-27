package repository

import (
	"be_project3team3/feature/cart/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Name     string
	Password string
	Phone    string
	Address  string
	ShopName string
	ImageUrl string
	Products []Product `gorm:"foreignKey:Id_user"`
	Carts    []Cart    `gorm:"foreignKey:Id_user"`
}

type Product struct {
	gorm.Model
	Id_user_seller uint
	Product_name   string
	Description    string
	Price          uint
	ImageUrl       string
	Stock          uint
	Category       string
	ShopName       string `gorm:"-:migration" gorm:"<-"`
	Carts          []Cart `gorm:"foreignKey:Id_product"`
}

type Cart struct {
	gorm.Model
	Id_user        uint
	Id_product     uint
	Product_name   string `gorm:"-:migration" gorm:"->"`
	Price          uint   `gorm:"-:migration" gorm:"->"`
	ImageUrl       string `gorm:"-:migration" gorm:"->"`
	Qty            uint
	Sub_total      string
	Notes          string
	ShopName       string `gorm:"-:migration" gorm:"->"`
	Id_user_seller uint
}

func FromDomain(du domain.Core) Cart {
	return Cart{
		Model:          gorm.Model{ID: du.ID},
		Id_user:        du.Id_user,
		Id_product:     du.Id_product,
		Product_name:   du.Product_name,
		Price:          du.Price,
		ImageUrl:       du.ImageUrl,
		Qty:            du.Qty,
		Sub_total:      du.Sub_total,
		Notes:          du.Notes,
		ShopName:       du.ShopName,
		Id_user_seller: du.Id_user_seller,
	}
}

func ToDomain(u Cart) domain.Core {
	return domain.Core{
		ID:             u.ID,
		Id_user:        u.Id_user,
		Id_product:     u.Id_product,
		Product_name:   u.Product_name,
		Price:          u.Price,
		ImageUrl:       u.ImageUrl,
		Qty:            u.Qty,
		Sub_total:      u.Sub_total,
		Notes:          u.Notes,
		ShopName:       u.ShopName,
		Id_user_seller: u.Id_user_seller,
	}
}

func ToDomainArray(au []Cart) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{
			ID:             val.ID,
			Id_user:        val.Id_user,
			Id_product:     val.Id_product,
			Product_name:   val.Product_name,
			Price:          val.Price,
			ImageUrl:       val.ImageUrl,
			Qty:            val.Qty,
			Sub_total:      val.Sub_total,
			Notes:          val.Notes,
			ShopName:       val.ShopName,
			Id_user_seller: val.Id_user_seller,
		})
	}
	return res
}
