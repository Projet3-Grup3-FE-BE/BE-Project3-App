package delivery

import (
	"be_project3team3/feature/cart/domain"
)

type PostDataCartFormat struct {
	Product_name   string `json:"product_name" form:"product_name"`
	Price          int    `json:"price" form:"price"`
	Qty            int    `json:"qty" form:"qty"`
	Sub_total      int    `json:"sub_total" form:"sub_total"`
	ImageUrl       string `json:"image_url" form:"image_url"`
	Notes          string `json:"notes" form:"notes"`
	Id_user        uint   `json:"id_user" form:"id_user"`
	Id_product     uint   `json:"id_product" form:"id_product"`
	Category       string `json:"category" form:"category"`
	ShopName       string `json:"shopname" form:"shopname"`
	Id_user_seller uint   `json:"id_user_seller" form:"id_user_seller"`
}

type UpdateCartFormat struct {
	ID        uint   `json:"id" form:"id"`
	Qty       int    `json:"qty" form:"qty"`
	Sub_total int    `json:"sub_total" form:"sub_total"`
	Notes     string `json:"notes" form:"notes"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case PostDataCartFormat:
		cnv := i.(PostDataCartFormat)
		return domain.Core{
			Product_name:   cnv.Product_name,
			Price:          cnv.Price,
			Qty:            cnv.Qty,
			Sub_total:      cnv.Sub_total,
			ImageUrl:       cnv.ImageUrl,
			Notes:          cnv.Notes,
			Id_user:        cnv.Id_user,
			Id_product:     cnv.Id_product,
			ShopName:       cnv.ShopName,
			Id_user_seller: cnv.Id_user_seller,
		}
	case UpdateCartFormat:
		cnv := i.(UpdateCartFormat)
		return domain.Core{
			ID:        cnv.ID,
			Qty:       cnv.Qty,
			Sub_total: cnv.Sub_total,
			Notes:     cnv.Notes,
		}
	}
	return domain.Core{}
}
