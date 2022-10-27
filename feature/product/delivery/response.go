package delivery

import (
	"be_project3team3/feature/product/domain"
	"log"
)

type productResponseFormat struct {
	ID             uint   `json:"id"`
	Product_Name   string `json:"product_name"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	Image_Url      string `json:"image_url"`
	Stock          int    `json:"stock"`
	Shop_Name      string `json:"shop_name"`
	Category       string `json:"category"`
	Id_User_Seller uint   `json:"id_user_seller"`
}

// type productCommentResponseFormat struct {
// 	ID             uint   `json:"id"`
// 	Product_Name   string `json:"product_name"`
// 	Description    string `json:"description"`
// 	Price          int    `json:"price"`
// 	Image_Url      string `json:"image_url"`
// 	Stock          int    `json:"stock"`
// 	Shop_Name      string `json:"shop_name"`
// 	Category       string `json:"category"`
// 	Id_User_Seller uint   `json:"id_user_seller"`
// }

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

type UploadResult struct {
	Path    string `json:"path" form:"path"`
	Content string `json:"content" form:"content"`
}

func SuccessDeleteResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "product":
		var product productResponseFormat
		cnv := core.(domain.Core)
		product = productResponseFormat{
			ID:             cnv.ID,
			Product_Name:   cnv.Product_Name,
			Description:    cnv.Description,
			Price:          cnv.Price,
			Image_Url:      cnv.Image_Url,
			Stock:          cnv.Stock,
			Shop_Name:      cnv.Shop_Name,
			Category:       cnv.Category,
			Id_User_Seller: cnv.Id_User_Seller,
		}
		res = product
	case "products":
		var arrproduct []productResponseFormat
		cnv := core.([]domain.Core)
		log.Println("\n\n isi res =", cnv, "\n\n")
		for _, val := range cnv {
			arrproduct = append(arrproduct,
				productResponseFormat{
					ID:             val.ID,
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
		res = arrproduct
	}
	return res
}

// func ToResponseproductComment(coreproduct interface{}, coreComment interface{}) interface{} {
// 	var res interface{}
// 	var product productCommentResponseFormat
// 	var arrComments []delComment.ResponseFormat
// 	cnvproduct := coreproduct.(domain.Core)
// 	cnvComments := coreComment.([]domComment.Core)

// 	log.Println("\n\n isi cnvproduct =", cnvproduct, "\n\n")
// 	log.Println("\n\n isi cnvComments =", cnvComments, "\n\n")

// 	for _, val := range cnvComments {
// 		arrComments = append(arrComments,
// 			delComment.ResponseFormat{
// 				ID:            val.ID,
// 				Name_User:     val.Name_User,
// 				Comment_Value: val.Comment_Value,
// 				IDUser:        val.IDUser,
// 				IDproduct:     val.IDproduct,
// 			})
// 	}

// 	product = productCommentResponseFormat{
// 		ID:        cnvproduct.ID,
// 		Name_User: cnvproduct.Name_User,
// 		Image_Url: cnvproduct.Image_Url,
// 		Content:   cnvproduct.Content,
// 		IDUser:    cnvproduct.IDUser,
// 		Comments:  arrComments,
// 	}
// 	res = product

// 	return res
// }
