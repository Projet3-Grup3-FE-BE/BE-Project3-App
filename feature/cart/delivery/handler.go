package delivery

import (
	"be_project3team3/config"
	"be_project3team3/feature/cart/domain"
	"be_project3team3/utils/jwt"
	"errors"
	"net/http"

	// "net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

type userHandler struct {
	srv domain.Service
}

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}

	e.POST("/carts", handler.AddDataCart(), middleware.JWT([]byte(key)))
	e.GET("/carts", handler.GetCart(), middleware.JWT([]byte(key)))
	e.PUT("/carts/:id", handler.UpdateCart(), middleware.JWT([]byte(key)))
	e.DELETE("/carts/:id", handler.DeleteByID(), middleware.JWT([]byte(key)))
}

// Done Add Cart
func (us *userHandler) AddDataCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input PostDataCartFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}
		input.Id_user = jwt.ExtractIdToken(c)
		cnv := ToDomain(input)
		res, err := us.srv.AddCart(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add cart", ToResponse(res, "reg")))
	}

}

// Done Updare cart
func (us *userHandler) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateCartFormat
		id, err := strconv.Atoi(c.Param("id"))
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind data"))
		}

		id_user := jwt.ExtractIdToken(c)
		input.ID = uint(id_user)
		cnv := ToDomain(input)
		res, err := us.srv.UpdateCart(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success update cart", ToResponse(res, "update")))
	}

}

// Done Delete
func (us *userHandler) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errors.New("cannot convert id")
		}
		_, err = us.srv.Delete(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request."))
		}
		return c.JSON(http.StatusOK, SuccessDelete("Success delete data."))
	}
}

// Done Get Cart
func (us *userHandler) GetCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := jwt.ExtractIdToken(c)
		res, err := us.srv.GetCart(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success show all data", ToResponseProduct(res, "sukses")))
	}
}
