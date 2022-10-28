package repository

import (
	"be_project3team3/feature/order/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ShippingName    string
	ShippingPhone   string
	ShippingAddress string
	ShopName        string
	GrossAmount     int
	OrderStatus     string
	SnapToken       string
	RedirectedUrl   string
	IdUserSeller    uint
	IdUserBuyer     uint
}

type Product struct {
	gorm.Model
	ProductName  string
	Description  string
	Price        int
	ImageUrl     string
	Stock        int
	ShopName     string
	Category     string
	IdUserSeller uint
}

func FromDomain(dom domain.Core) Order {
	return Order{
		Model:           gorm.Model{ID: dom.ID},
		ShippingName:    dom.ShippingName,
		ShippingPhone:   dom.ShippingPhone,
		ShippingAddress: dom.ShippingAddress,
		ShopName:        dom.ShopName,
		GrossAmount:     dom.GrossAmount,
		OrderStatus:     dom.OrderStatus,
		SnapToken:       dom.SnapToken,
		RedirectedUrl:   dom.RedirectedUrl,
		IdUserSeller:    dom.IdUserSeller,
		IdUserBuyer:     dom.IdUserBuyer,
	}
}

func ToDomain(p Order) domain.Core {
	return domain.Core{
		ID:              p.ID,
		ShippingName:    p.ShippingName,
		ShippingPhone:   p.ShippingPhone,
		ShippingAddress: p.ShippingAddress,
		ShopName:        p.ShopName,
		GrossAmount:     p.GrossAmount,
		OrderStatus:     p.OrderStatus,
		SnapToken:       p.SnapToken,
		RedirectedUrl:   p.RedirectedUrl,
		IdUserSeller:    p.IdUserSeller,
		IdUserBuyer:     p.IdUserBuyer,
	}
}

func ToDomainArray(arrproduct []Order) []domain.Core {
	var res []domain.Core
	for _, val := range arrproduct {
		res = append(res, domain.Core{ID: val.ID,
			ShippingName:    val.ShippingName,
			ShippingPhone:   val.ShippingPhone,
			ShippingAddress: val.ShippingAddress,
			ShopName:        val.ShopName,
			GrossAmount:     val.GrossAmount,
			OrderStatus:     val.OrderStatus,
			IdUserSeller:    val.IdUserSeller,
			IdUserBuyer:     val.IdUserBuyer,
		})
	}

	return res
}

// func ToDomainproductComment(pproduct product, arrComment []repoComment.Comment) []domain.Core {
// 	var res []domain.Core
// 	for _, val := range arrproduct {
// 		res = append(res, domain.Core{ID: val.ID,
// 			Name_User: val.Name_User,
// 			Image_Url: val.Image_Url,
// 			Content:   val.Content,
// 			IDUser:    val.IDUser,
// 		})
// 	}

// 	return res
// }
