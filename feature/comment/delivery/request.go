package delivery

import (
	"be_project3team3/feature/comment/domain"
)

type InsertRequestFormat struct {
	Comment_Value string `json:"comment_value" form:"comment_value"`
	IDPosting     uint   `json:"id_posting" form:"id_posting"`
}
type RequestFormat struct {
	ID            uint   `json:"id" form:"id"`
	Name_User     string `json:"name_user" form:"name_user"`
	Comment_Value string `json:"comment_value" form:"comment_value"`
	IDUser        uint   `json:"id_user" form:"id_user"`
	IDPosting     uint   `json:"id_posting" form:"id_posting"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case InsertRequestFormat:
		cnv := i.(InsertRequestFormat)
		return domain.Core{
			Comment_Value: cnv.Comment_Value,
			IDPosting:     cnv.IDPosting,
		}
	case RequestFormat:
		cnv := i.(RequestFormat)
		return domain.Core{
			ID:            cnv.ID,
			Name_User:     cnv.Name_User,
			Comment_Value: cnv.Comment_Value,
			IDUser:        cnv.IDUser,
			IDPosting:     cnv.IDPosting,
		}
	}
	return domain.Core{}
}
