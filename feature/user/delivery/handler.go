package delivery

import (
	"be_project3team3/config"
	"be_project3team3/feature/user/domain"

	// "log"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
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
	// e.GET("/users", handler.ShowAllUser())
	// e.GET("/users/:email", handler.Profile(), middleware.JWT([]byte(key)))
	// e.PUT("/users", handler.EditProfile(), middleware.JWT([]byte(key)))
	// e.DELETE("/users", handler.DeleteUser(), middleware.JWT([]byte(key)))
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
