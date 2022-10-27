package repository

import (
	"be_project3team3/feature/product/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Product_Name   string
	Description    string
	Price          int
	Image_Url      string
	Stock          int
	Shop_Name      string
	Category       string
	Id_User_Seller uint
}

func FromDomain(dom domain.Core) Product {
	return Product{
		Model:          gorm.Model{ID: dom.ID},
		Product_Name:   dom.Product_Name,
		Description:    dom.Description,
		Price:          dom.Price,
		Image_Url:      dom.Image_Url,
		Stock:          dom.Stock,
		Shop_Name:      dom.Shop_Name,
		Category:       dom.Category,
		Id_User_Seller: dom.Id_User_Seller,
	}
}

func ToDomain(p Product) domain.Core {
	return domain.Core{
		ID:             p.ID,
		Product_Name:   p.Product_Name,
		Description:    p.Description,
		Price:          p.Price,
		Image_Url:      p.Image_Url,
		Stock:          p.Stock,
		Shop_Name:      p.Shop_Name,
		Category:       p.Category,
		Id_User_Seller: p.Id_User_Seller,
	}
}

func ToDomainArray(arrproduct []Product) []domain.Core {
	var res []domain.Core
	for _, val := range arrproduct {
		res = append(res, domain.Core{ID: val.ID,
			Product_Name:   val.Product_Name,
			Description:    val.Description,
			Price:          val.Price,
			Image_Url:      val.Image_Url,
			Stock:          val.Stock,
			Shop_Name:      val.Shop_Name,
			Category:       val.Category,
			Id_User_Seller: val.Id_User_Seller,
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
