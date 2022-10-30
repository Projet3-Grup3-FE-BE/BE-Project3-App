package delivery

import (
	"be_project3team3/feature/order/domain"
)

type OrderInsertRequestFormat struct {
	ShippingName    string `json:"shipping_name" form:"shipping_name"`
	ShippingPhone   string `json:"shipping_phone" form:"shipping_phone"`
	ShippingAddress string `json:"shipping_address" form:"shipping_address"`
	ShopName        string `json:"shop_name" form:"shop_name"`
	GrossAmount     int    `json:"gross_amount" form:"gross_amount"`
	IdUserSeller    uint   `json:"id_user_seller" form:"id_user_seller"`
}

type OrderRequestFormat struct {
	ID              uint   `json:"id" form:"id"`
	ShippingName    string `json:"shipping_name" form:"shipping_name"`
	ShippingPhone   string `json:"shipping_phone" form:"shipping_phone"`
	ShippingAddress string `json:"shipping_address" form:"shipping_address"`
	ShopName        string `json:"shop_name" form:"shop_name"`
	GrossAmount     int    `json:"gross_amount" form:"gross_amount"`
	OrderStatus     string `json:"order_status" form:"order_status"`
	SnapToken       string `json:"snap_token" form:"snap_token"`
	RedirectedUrl   string `json:"redirect_url" form:"redirect_url"`
	IdUserSeller    uint   `json:"id_user_seller" form:"id_user_seller"`
	IdUserBuyer     uint   `json:"id_user_buyer" form:"id_user_buyer"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case OrderInsertRequestFormat:
		cnv := i.(OrderInsertRequestFormat)
		return domain.Core{
			ShippingName:    cnv.ShippingName,
			ShippingPhone:   cnv.ShippingPhone,
			ShippingAddress: cnv.ShippingAddress,
			ShopName:        cnv.ShopName,
			GrossAmount:     cnv.GrossAmount,
			IdUserSeller:    cnv.IdUserSeller,
			IdUserBuyer:     cnv.IdUserSeller,
		}
	case OrderRequestFormat:
		cnv := i.(OrderRequestFormat)
		return domain.Core{
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
			IdUserBuyer:     cnv.IdUserSeller,
		}
	}
	return domain.Core{}
}
