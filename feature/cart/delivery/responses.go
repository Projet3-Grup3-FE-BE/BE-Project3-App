package delivery

import (
	"be_project3team3/feature/cart/domain"
)

func SuccessDelete(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}
func SuccessLogin(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

type PostCartRespons struct {
	ID             uint   `json:"id"`
	Product_name   string `json:"product_name"`
	Price          int    `json:"price"`
	Qty            int    `json:"qty"`
	Sub_total      int    `json:"sub_total"`
	ImageUrl       string `json:"image_url"`
	Notes          string `json:"notes"`
	Id_user_buyer  uint   `json:"id_user_buyer"`
	Id_product     uint   `json:"id_product"`
	Category       string `json:"category"`
	ShopName       string `json:"shopname"`
	Id_user_seller uint   `json:"id_user_seller"`
}

type GetCartIDRespons struct {
	ID             uint   `json:"id"`
	Product_name   string `json:"product_name"`
	Price          int    `json:"price"`
	Qty            int    `json:"qty"`
	Sub_total      int    `json:"sub_total"`
	ImageUrl       string `json:"image_url"`
	Notes          string `json:"notes"`
	Category       string `json:"category"`
	ShopName       string `json:"shopname"`
	Id_product     uint   `json:"id_product"`
	Id_user_seller uint   `json:"id_user_seller"`
	Id_user        uint   `json:"id_user"`
}

type UpdateCartRespons struct {
	ID             uint   `json:"id"`
	Product_name   string `json:"product_name"`
	Price          int    `json:"price"`
	Qty            int    `json:"qty"`
	Sub_total      int    `json:"sub_total"`
	ImageUrl       string `json:"image_url"`
	Notes          string `json:"notes"`
	Id_user_buyer  uint   `json:"id_user_buyer"`
	Id_product     uint   `json:"id_product"`
	Category       string `json:"category"`
	ShopName       string `json:"shopname"`
	Id_user_seller uint   `json:"id_user_seller"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = PostCartRespons{
			ID:           cnv.ID,
			Product_name: cnv.Product_name,
			Price:        cnv.Price,
			Qty:          cnv.Qty,
			Sub_total:    cnv.Sub_total,
			ImageUrl:     cnv.ImageUrl,
			Notes:        cnv.Notes,
			// Category:  cnv.Category,
			ShopName:       cnv.ShopName,
			Id_product:     cnv.Id_product,
			Id_user_seller: cnv.Id_user_seller,
			// Id_user_buyer: cnv.Id_user_buyer,
		}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateCartRespons{
			ID:           cnv.ID,
			Product_name: cnv.Product_name,
			Price:        cnv.Price,
			Qty:          cnv.Qty,
			Sub_total:    cnv.Sub_total,
			ImageUrl:     cnv.ImageUrl,
			Notes:        cnv.Notes,
			// Category:  cnv.Category,
			ShopName:       cnv.ShopName,
			Id_product:     cnv.Id_product,
			Id_user_seller: cnv.Id_user_seller,
			// Id_user_buyer: cnv.Id_user_buyer,
		}
	}
	return res
}

func ToResponseProduct(core interface{}, code string) interface{} {
	var res interface{}
	var arr []GetCartIDRespons
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, GetCartIDRespons{
			ID:           cnv.ID,
			Product_name: cnv.Product_name,
			Price:        cnv.Price,
			Qty:          cnv.Qty,
			Sub_total:    cnv.Sub_total,
			ImageUrl:     cnv.ImageUrl,
			Notes:        cnv.Notes,
			// Category:  cnv.Category,
			ShopName:       cnv.ShopName,
			Id_product:     cnv.Id_product,
			Id_user_seller: cnv.Id_user_seller,
			Id_user:        cnv.Id_user,
		})
	}
	res = arr
	return res
}
