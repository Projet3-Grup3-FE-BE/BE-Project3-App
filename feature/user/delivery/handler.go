package delivery

import (
	"be_project3team3/config"
	"be_project3team3/feature/user/domain"
	jwt "be_project3team3/utils/jwt"

	// "log"
	"net/http"

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

	e.POST("/register", handler.Register())
	e.POST("/login", handler.LoginUser())
	e.PUT("/users", handler.UpdateUser(), middleware.JWT([]byte("R4hs!!a@")))
	e.DELETE("/users", handler.DeleteByID(), middleware.JWT([]byte("R4hs!!a@")))
	// e.GET("/users", handler.ShowAllUser())
	// e.GET("/users/:email", handler.Profile(), middleware.JWT([]byte(key)))

}

// registrasi add user
func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (us *userHandler) LoginUser() echo.HandlerFunc {
	//autentikasi user login
	return func(c echo.Context) error {
		var resQry LoginFormat
		if err := c.Bind(&resQry); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToDomain(resQry)
		res, err := us.srv.LoginUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		token := us.srv.GenerateToken(res.ID)
		return c.JSON(http.StatusCreated, SuccessLogin("berhasil login", token, ToResponse(res, "login")))
	}
}

func (us *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		input.Email = c.FormValue("email")
		input.Password = c.FormValue("password")
		input.Name = c.FormValue("name")
		input.Phone = c.FormValue("phone")
		input.Bio = c.FormValue("bio")
		input.Gender = c.FormValue("gender")
		input.Location = c.FormValue("location")

		input.ID = jwt.ExtractIdToken(c)
		if input.ID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}

		cnv := ToDomain(input)
		res, err := us.srv.UpdateProfile(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "upd")))
	}

}

func (us *userHandler) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := jwt.ExtractIdToken(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "cannot validate token",
			})
		}
		err := us.srv.Delete(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessDelete("success delete user"))
	}
}
