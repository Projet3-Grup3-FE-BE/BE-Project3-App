package delivery

import (
	"be_project3team3/feature/product/domain"
)

type productInsertRequestFormat struct {
	Product_Name   string `json:"product_name" form:"product_name"`
	Description    string `json:"description" form:"description"`
	Price          int    `json:"price" form:"price"`
	Image_Url      string `json:"image_url" form:"image_url"`
	Stock          int    `json:"stock" form:"stock"`
	Shop_Name      string `json:"shop_name" form:"shop_name"`
	Category       string `json:"category" form:"category"`
	Id_User_Seller uint   `json:"id_user_seller" form:"id_user_seller"`
}

type productRequestFormat struct {
	ID             uint   `json:"id" form:"id"`
	Product_Name   string `json:"product_name" form:"product_name"`
	Description    string `json:"description" form:"description"`
	Price          int    `json:"price" form:"price"`
	Image_Url      string `json:"image_url" form:"image_url"`
	Stock          int    `json:"stock" form:"stock"`
	Shop_Name      string `json:"shop_name" form:"shop_name"`
	Category       string `json:"category" form:"category"`
	Id_User_Seller uint   `json:"id_user_seller" form:"id_user_seller"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case productInsertRequestFormat:
		cnv := i.(productInsertRequestFormat)
		return domain.Core{
			Product_Name:   cnv.Product_Name,
			Description:    cnv.Description,
			Price:          cnv.Price,
			Image_Url:      cnv.Image_Url,
			Stock:          cnv.Stock,
			Shop_Name:      cnv.Shop_Name,
			Category:       cnv.Category,
			Id_User_Seller: cnv.Id_User_Seller,
		}
	case productRequestFormat:
		cnv := i.(productRequestFormat)
		return domain.Core{
			ID:             cnv.ID,
			Product_Name:   cnv.Product_Name,
			Description:    cnv.Description,
			Price:          cnv.Price,
			Image_Url:      cnv.Image_Url,
			Stock:          cnv.Stock,
			Category:       cnv.Category,
			Id_User_Seller: cnv.Id_User_Seller,
		}
	}
	return domain.Core{}
}
