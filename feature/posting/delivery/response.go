package delivery

import (
	delComment "be_project3team3/feature/comment/delivery"
	domComment "be_project3team3/feature/comment/domain"
	"be_project3team3/feature/posting/domain"
	"log"
)

type PostingResponseFormat struct {
	ID        uint   `json:"id"`
	Name_User string `json:"name_user"`
	Image_Url string `json:"image_url"`
	Content   string `json:"content"`
	IDUser    uint   `json:"id_user"`
}
type PostingCommentResponseFormat struct {
	ID        uint                        `json:"id"`
	Name_User string                      `json:"name_user"`
	Image_Url string                      `json:"image_url"`
	Content   string                      `json:"content"`
	IDUser    uint                        `json:"id_user"`
	Comments  []delComment.ResponseFormat `json:"comments"`
}

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
	case "posting":
		var posting PostingResponseFormat
		cnv := core.(domain.Core)
		posting = PostingResponseFormat{
			ID:        cnv.ID,
			Name_User: cnv.Name_User,
			Image_Url: cnv.Image_Url,
			Content:   cnv.Content,
			IDUser:    cnv.IDUser,
		}
		res = posting
	case "postings":
		var arrPosting []PostingResponseFormat
		cnv := core.([]domain.Core)
		log.Println("\n\n isi res =", cnv, "\n\n")
		for _, val := range cnv {
			arrPosting = append(arrPosting,
				PostingResponseFormat{
					ID:        val.ID,
					Name_User: val.Name_User,
					Image_Url: val.Image_Url,
					Content:   val.Content,
					IDUser:    val.IDUser,
				})
		}
		res = arrPosting
	}
	return res
}

func ToResponsePostingComment(corePosting interface{}, coreComment interface{}) interface{} {
	var res interface{}
	var posting PostingCommentResponseFormat
	var arrComments []delComment.ResponseFormat
	cnvPosting := corePosting.(domain.Core)
	cnvComments := coreComment.([]domComment.Core)

	log.Println("\n\n isi cnvPosting =", cnvPosting, "\n\n")
	log.Println("\n\n isi cnvComments =", cnvComments, "\n\n")

	for _, val := range cnvComments {
		arrComments = append(arrComments,
			delComment.ResponseFormat{
				ID:            val.ID,
				Name_User:     val.Name_User,
				Comment_Value: val.Comment_Value,
				IDUser:        val.IDUser,
				IDPosting:     val.IDPosting,
			})
	}

	posting = PostingCommentResponseFormat{
		ID:        cnvPosting.ID,
		Name_User: cnvPosting.Name_User,
		Image_Url: cnvPosting.Image_Url,
		Content:   cnvPosting.Content,
		IDUser:    cnvPosting.IDUser,
		Comments:  arrComments,
	}
	res = posting

	return res
}
