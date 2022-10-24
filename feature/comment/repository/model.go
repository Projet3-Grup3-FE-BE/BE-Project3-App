package repository

import (
	"be_project3team3/feature/comment/domain"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID            uint
	Name_User     string
	Comment_Value string
	IDUser        uint
	IDPosting     uint
}

func FromDomain(dom domain.Core) Comment {
	return Comment{
		Model:         gorm.Model{ID: dom.ID},
		ID:            dom.ID,
		Name_User:     dom.Name_User,
		Comment_Value: dom.Comment_Value,
		IDUser:        dom.IDUser,
		IDPosting:     dom.IDPosting,
	}
}

func ToDomain(c Comment) domain.Core {
	return domain.Core{
		ID:            c.ID,
		Name_User:     c.Name_User,
		Comment_Value: c.Comment_Value,
		IDUser:        c.IDUser,
		IDPosting:     c.IDPosting,
	}
}

func ToDomainArray(arrComment []Comment) []domain.Core {
	var res []domain.Core
	for _, val := range arrComment {
		res = append(res, domain.Core{ID: val.ID,
			Name_User:     val.Name_User,
			Comment_Value: val.Comment_Value,
			IDUser:        val.IDUser,
			IDPosting:     val.IDPosting,
		})
	}

	return res
}
