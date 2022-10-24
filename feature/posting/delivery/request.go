package delivery

import (
	"be_project3team3/feature/posting/domain"
)

type PostingInsertRequestFormat struct {
	Name_User string `json:"name_user" form:"name_user"`
	Image_Url string `json:"image_url" form:"image_url"`
	Content   string `json:"content" form:"content"`
}

type PostingRequestFormat struct {
	ID        uint   `json:"id" form:"id"`
	Name_User string `json:"name_user" form:"name_user"`
	Image_Url string `json:"image_url" form:"image_url"`
	Content   string `json:"content" form:"content"`
	IDUser    uint   `json:"id_user" form:"id_user"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case PostingInsertRequestFormat:
		cnv := i.(PostingInsertRequestFormat)
		return domain.Core{
			Name_User: cnv.Name_User,
			Image_Url: cnv.Image_Url,
			Content:   cnv.Content,
		}
	case PostingRequestFormat:
		cnv := i.(PostingRequestFormat)
		return domain.Core{
			ID:        cnv.ID,
			Name_User: cnv.Name_User,
			Image_Url: cnv.Image_Url,
			Content:   cnv.Content,
			IDUser:    cnv.IDUser,
		}
	}
	return domain.Core{}
}
