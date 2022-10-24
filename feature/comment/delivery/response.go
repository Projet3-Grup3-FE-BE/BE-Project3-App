package delivery

import (
	"be_project3team3/feature/comment/domain"
	"log"
)

type ResponseFormat struct {
	ID            uint   `json:"id"`
	Name_User     string `json:"name_user"`
	Comment_Value string `json:"comment_value"`
	IDUser        uint   `json:"id_user"`
	IDPosting     uint   `json:"id_posting"`
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
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
	case "comment":
		var comment ResponseFormat
		cnv := core.(domain.Core)
		comment = ResponseFormat{
			ID:            cnv.ID,
			Name_User:     cnv.Name_User,
			Comment_Value: cnv.Comment_Value,
			IDUser:        cnv.IDUser,
			IDPosting:     cnv.IDPosting,
		}
		res = comment
	case "comments":
		var arr []ResponseFormat
		cnv := core.([]domain.Core)
		log.Println("\n\n isi res =", cnv, "\n\n")
		for _, val := range cnv {
			arr = append(arr,
				ResponseFormat{
					ID:            val.ID,
					Name_User:     val.Name_User,
					Comment_Value: val.Comment_Value,
					IDUser:        val.IDUser,
					IDPosting:     val.IDPosting,
				})
		}
		res = arr
	case "postingAllComments":
		// pending, nunggu comment dibuat

		// var arrPosting []PostingResponseFormat
		// cnv := core.([]domain.Core)
		// log.Println("\n\n isi res =", cnv, "\n\n")
		// for _, val := range cnv {
		// 	arrPosting = append(arrPosting,
		// 		PostingResponseFormat{
		// 			ID:        val.ID,
		// 			Name_User: val.Name_User,
		// 			Image_Url: val.Image_Url,
		// 			Content:   val.Content,
		// 			IDUser:    val.IDUser,
		// 		})
		// }
		// res = arrPosting
	}

	return res
}
