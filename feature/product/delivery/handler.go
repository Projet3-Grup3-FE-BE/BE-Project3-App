package delivery

import (
	"be_project3team3/config"
	"be_project3team3/feature/product/domain"
	"be_project3team3/utils/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

type productHandler struct {
	srv domain.ServiceInterface
}

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func New(e *echo.Echo, srv domain.ServiceInterface) {
	handler := productHandler{srv: srv}
	e.POST("/products", handler.Addproduct(), middleware.JWT([]byte(key))) //jangan lupa dikasih jwt soalnya habis di hapus buat coba
	e.GET("/products", handler.GetAllproduct())
	e.GET("/products/:id", handler.Getproduct())
	// e.GET("/products/:id/comments", handler.GetproductAllComment())
	e.PUT("/products/:id", handler.Updateproduct(), middleware.JWT([]byte(key)))
	e.DELETE("/products/:id", handler.Deleteproduct(), middleware.JWT([]byte(key)))
}

func (us *productHandler) GetAllproduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		category := c.QueryParam("category")
		id_user_seller := c.QueryParam("id_user_seller")
		res, err := us.srv.GetAll(id_user_seller, category)
		//log.Println("\n\n\n res GET ALL =", res, "\n\n")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get products.", ToResponse(res, "products")))
	}
}

func (us *productHandler) Getproduct() echo.HandlerFunc {
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

		return c.JSON(http.StatusOK, SuccessResponse("Success get product.", ToResponse(res, "product")))
	}
}

// func (us *productHandler) GetproductAllComment() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// Check authorized request atau tidak dgn token

// 		paramID := c.Param("id")
// 		if paramID == "" {
// 			return c.JSON(http.StatusInternalServerError, FailResponse("Failed. Id empty or not found."))
// 		}
// 		resproduct, resComments, err := us.srv.GetproductAllComment(paramID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
// 		}

// 		return c.JSON(http.StatusOK, SuccessResponse("Success get product.", ToResponseproductComment(resproduct, resComments)))
// 	}
// }

func (us *productHandler) Addproduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token

		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}
		var input productInsertRequestFormat
		// uploader = NewUploader()
		// //content := c.FormValue("content")
		// isSuccess := true
		// file, er := c.FormFile("file")
		// if er != nil {
		// 	isSuccess = false
		// } else {
		// 	src, err := file.Open()
		// 	if err != nil {
		// 		isSuccess = false
		// 	} else {
		// 		resFile, err := upload(c, file.Filename, src)
		// 		if err != nil {
		// 			return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
		// 		}
		// 		input.Image_Url = resFile
		// 	}
		// 	defer src.Close()
		// }
		isSuccess := true
		input.Image_Url = "https://ds393qgzrxwzn.cloudfront.net/resize/m720x480/cat1/img/images/0/ISzO90zLnp.jpg"
		if isSuccess {
			if err := c.Bind(&input); err != nil {
				log.Println("Error Bind = ", err.Error())
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			log.Println("\n\n\n input product handler : ", input, "\n\n\n")
			id := jwt.ExtractIdToken(c)
			cnv := ToDomain(input)
			cnv.Id_User_Seller = id
			log.Println("\n\n\n input product handler cnv : ", cnv, "\n\n\n")

			res, err := us.srv.Insert(cnv, c)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusCreated, SuccessResponse("Success add product.", ToResponse(res, "product")))

		}

		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
	}

}

func (us *productHandler) Updateproduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponse(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		var input productRequestFormat
		paramID := c.Param("id")
		// uploader = NewUploader()

		// isSuccess := true
		// file, er := c.FormFile("file")
		// if er != nil {
		// 	isSuccess = false
		// } else {
		// 	src, err := file.Open()
		// 	if err != nil {
		// 		isSuccess = false
		// 	} else {
		// 		resFile, err := upload(c, file.Filename, src)
		// 		if err != nil {
		// 			return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
		// 		}
		// 		input.Image_Url = resFile
		// 	}
		// 	defer src.Close()
		// }
		isSuccess := true
		input.Image_Url = "https://ds393qgzrxwzn.cloudfront.net/resize/m720x480/cat1/img/images/0/ISzO90zLnp.jpg"
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

			return c.JSON(http.StatusCreated, SuccessResponse("Success update product.", ToResponse(res, "product")))

		}
		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
	}
}

func (us *productHandler) Deleteproduct() echo.HandlerFunc {
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
					return c.JSON(http.StatusCreated, SuccessDeleteResponse("Success delete product."))
				}
			}
		}

	}
}

// var uploader *s3manager.Uploader

// func NewUploader() *s3manager.Uploader {
// 	s3Config := &aws.Config{
// 		Region:      aws.String("ap-southeast-1"),
// 		Credentials: credentials.NewStaticCredentials("AKIARQA2KZ55LJN2AW7L", "zsp4Vtew2D/dTYjHQj48WNmSJUP/WJ3m2wm66qIm", ""),
// 	}
// 	s3Session := session.New(s3Config)
// 	uploader := s3manager.NewUploader(s3Session)
// 	return uploader
// }

// func upload(c echo.Context, filename string, src multipart.File) (string, error) {
// 	logger := c.Logger()
// 	log.Println("uploading")

// 	upInput := &s3manager.UploadInput{
// 		Bucket: aws.String("projectalta"), // bucket's name
// 		Key:    aws.String(filename),      // files destination location
// 		Body:   src,                       // content of the file
// 		//ContentType: aws.String("image/jpg"),   // content type
// 	}
// 	res, err := uploader.UploadWithContext(context.Background(), upInput)
// 	if err != nil {
// 		logger.Fatal(err)
// 		return "", err
// 	}
// 	return res.Location, nil
// }
