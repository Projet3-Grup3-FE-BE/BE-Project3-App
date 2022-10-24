package delivery

import (
	"be_project3team3/config"
	"be_project3team3/feature/posting/domain"
	"be_project3team3/utils/jwt"
	"context"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

type postingHandler struct {
	srv domain.ServiceInterface
}

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func New(e *echo.Echo, srv domain.ServiceInterface) {
	handler := postingHandler{srv: srv}
	e.POST("/postings", handler.AddPosting(), middleware.JWT([]byte(key))) //jangan lupa dikasih jwt soalnya habis di hapus buat coba
	e.GET("/postings", handler.GetAllPosting())
	e.GET("/postings/:id", handler.GetPosting())
	e.GET("/postings/:id/comments", handler.GetPostingAllComment())
	e.PUT("/postings/:id", handler.UpdatePosting(), middleware.JWT([]byte(key)))
	e.DELETE("/postings/:id", handler.DeletePosting(), middleware.JWT([]byte(key)))
}

func (us *postingHandler) GetAllPosting() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := us.srv.GetAll()
		//log.Println("\n\n\n res GET ALL =", res, "\n\n")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get postings.", ToResponse(res, "postings")))
	}
}

func (us *postingHandler) GetPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		// err := us.srv.IsAuthorized(c)
		// if err != nil {
		// 	return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		// } else {
		// 	log.Println("Authorized request.")
		// }

		paramID := c.Param("id")
		res, err := us.srv.Get(paramID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get posting.", ToResponse(res, "posting")))
	}
}

func (us *postingHandler) GetPostingAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token

		paramID := c.Param("id")
		if paramID == "" {
			return c.JSON(http.StatusInternalServerError, FailResponse("Failed. Id empty or not found."))
		}
		resPosting, resComments, err := us.srv.GetPostingAllComment(paramID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get posting.", ToResponsePostingComment(resPosting, resComments)))
	}
}

func (us *postingHandler) AddPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token

		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}
		uploader = NewUploader()
		var input PostingInsertRequestFormat
		//content := c.FormValue("content")
		isSuccess := true
		file, er := c.FormFile("file")
		if er != nil {
			isSuccess = false
		} else {
			src, err := file.Open()
			if err != nil {
				isSuccess = false
			} else {
				resFile, err := upload(c, file.Filename, src)
				if err != nil {
					return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
				}
				input.Image_Url = resFile
			}
			defer src.Close()
		}
		if isSuccess {
			if err := c.Bind(&input); err != nil {
				log.Println("Error Bind = ", err.Error())
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			log.Println("\n\n\n input posting handler : ", input, "\n\n\n")
			id := jwt.ExtractIdToken(c)
			cnv := ToDomain(input)
			cnv.IDUser = id
			log.Println("\n\n\n input posting handler cnv : ", cnv, "\n\n\n")

			res, err := us.srv.Insert(cnv, c)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusCreated, SuccessResponse("Success add posting.", ToResponse(res, "posting")))

		}

		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
	}

}

func (us *postingHandler) UpdatePosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		var input PostingRequestFormat
		paramID := c.Param("id")
		uploader = NewUploader()

		isSuccess := true
		file, er := c.FormFile("file")
		if er != nil {
			isSuccess = false
		} else {
			src, err := file.Open()
			if err != nil {
				isSuccess = false
			} else {
				resFile, err := upload(c, file.Filename, src)
				if err != nil {
					return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
				}
				input.Image_Url = resFile
			}
			defer src.Close()
		}
		if isSuccess {

			if err := c.Bind(&input); err != nil {
				log.Println("Error Bind = ", err.Error())
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			log.Println("\n\n\nid handler : ", paramID)
			log.Println("\n\n\n input handler : ", input)

			cnv := ToDomain(input)
			res, err := us.srv.Update(cnv, paramID, c)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}

			return c.JSON(http.StatusCreated, SuccessResponse("Success update posting.", ToResponse(res, "posting")))

		}
		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
	}
}

func (us *postingHandler) DeletePosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		paramID := strings.TrimSpace(c.Param("id"))
		// validasi not empty id
		if paramID == "" {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		} else {
			res, err := us.srv.Get(paramID)
			log.Println("res data :", res)
			// validasi get data by id
			if err != nil {
				log.Println("Error get data. Error :", err.Error())
				return c.JSON(http.StatusInternalServerError, FailResponse("Data not found"))
			} else {
				res2, err2 := us.srv.Delete(paramID)
				log.Println("res2 data :", res2)
				if err2 != nil {
					return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
				} else {
					return c.JSON(http.StatusCreated, SuccessDeleteResponse("Success delete posting."))
				}
			}
		}

	}
}

var uploader *s3manager.Uploader

func NewUploader() *s3manager.Uploader {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIARQA2KZ55LJN2AW7L", "zsp4Vtew2D/dTYjHQj48WNmSJUP/WJ3m2wm66qIm", ""),
	}
	s3Session := session.New(s3Config)
	uploader := s3manager.NewUploader(s3Session)
	return uploader
}

func upload(c echo.Context, filename string, src multipart.File) (string, error) {
	logger := c.Logger()
	log.Println("uploading")

	upInput := &s3manager.UploadInput{
		Bucket: aws.String("projectalta"), // bucket's name
		Key:    aws.String(filename),      // files destination location
		Body:   src,                       // content of the file
		//ContentType: aws.String("image/jpg"),   // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), upInput)
	if err != nil {
		logger.Fatal(err)
		return "", err
	}
	return res.Location, nil
}
