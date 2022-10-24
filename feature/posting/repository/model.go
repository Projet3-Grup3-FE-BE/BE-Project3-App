package repository

import (
	"be_project3team3/feature/posting/domain"

	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	ID        uint
	Name_User string
	Image_Url string
	Content   string
	IDUser    uint
}

func FromDomain(dom domain.Core) Posting {
	return Posting{
		Model:     gorm.Model{ID: dom.ID},
		ID:        dom.ID,
		Name_User: dom.Name_User,
		Image_Url: dom.Image_Url,
		Content:   dom.Content,
		IDUser:    dom.IDUser,
	}
}

func ToDomain(p Posting) domain.Core {
	return domain.Core{
		ID:        p.ID,
		Name_User: p.Name_User,
		Image_Url: p.Image_Url,
		Content:   p.Content,
		IDUser:    p.IDUser,
	}
}

func ToDomainArray(arrPosting []Posting) []domain.Core {
	var res []domain.Core
	for _, val := range arrPosting {
		res = append(res, domain.Core{ID: val.ID,
			Name_User: val.Name_User,
			Image_Url: val.Image_Url,
			Content:   val.Content,
			IDUser:    val.IDUser,
		})
	}

	return res
}

// func ToDomainPostingComment(pPosting Posting, arrComment []repoComment.Comment) []domain.Core {
// 	var res []domain.Core
// 	for _, val := range arrPosting {
// 		res = append(res, domain.Core{ID: val.ID,
// 			Name_User: val.Name_User,
// 			Image_Url: val.Image_Url,
// 			Content:   val.Content,
// 			IDUser:    val.IDUser,
// 		})
// 	}

// 	return res
// }
