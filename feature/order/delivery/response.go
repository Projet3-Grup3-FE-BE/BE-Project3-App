package delivery

import (
	"be_project3team3/feature/order/domain"
	"log"
)

type OrderResponseFormat struct {
	ID              uint   `json:"id" form:"id"`
	ShippingName    string `json:"shipping_name"`
	ShippingPhone   string `json:"shipping_phone"`
	ShippingAddress string `json:"shipping_address"`
	ShopName        string `json:"shop_name"`
	GrossAmount     int    `json:"gross_amount"`
	OrderStatus     string `json:"order_status"`
	SnapToken       string `json:"snap_token" `
	RedirectedUrl   string `json:"redirect_url"`
	IdUserSeller    uint   `json:"id_user_seller"`
	IdUserBuyer     uint   `json:"id_user_buyer"`
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
	case "order":
		var Order OrderResponseFormat
		cnv := core.(domain.Core)
		Order = OrderResponseFormat{
			ID:              cnv.ID,
			ShippingName:    cnv.ShippingName,
			ShippingPhone:   cnv.ShippingPhone,
			ShippingAddress: cnv.ShippingAddress,
			ShopName:        cnv.ShopName,
			GrossAmount:     cnv.GrossAmount,
			OrderStatus:     cnv.OrderStatus,
			SnapToken:       cnv.SnapToken,
			RedirectedUrl:   cnv.RedirectedUrl,
			IdUserSeller:    cnv.IdUserSeller,
			IdUserBuyer:     cnv.IdUserBuyer,
		}
		res = Order
	case "orders":
		var arrOrder []OrderResponseFormat
		cnv := core.([]domain.Core)
		log.Println("\n\n isi res =", cnv, "\n\n")
		for _, val := range cnv {
			arrOrder = append(arrOrder,
				OrderResponseFormat{
					ID:              val.ID,
					ShippingName:    val.ShippingName,
					ShippingPhone:   val.ShippingPhone,
					ShippingAddress: val.ShippingAddress,
					ShopName:        val.ShopName,
					GrossAmount:     val.GrossAmount,
					OrderStatus:     val.OrderStatus,
					SnapToken:       val.SnapToken,
					RedirectedUrl:   val.RedirectedUrl,
					IdUserSeller:    val.IdUserSeller,
					IdUserBuyer:     val.IdUserBuyer,
				})
		}
		res = arrOrder
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
