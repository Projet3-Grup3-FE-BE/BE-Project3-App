package delivery

import (
	"be_project3team3/config"
	"be_project3team3/feature/user/domain"
	"log"
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
	e.POST("/login", handler.Login())
	e.GET("/users", handler.ShowAllUser())
	e.GET("/users/:email", handler.Profile(), middleware.JWT([]byte(key)))
	e.PUT("/users", handler.EditProfile(), middleware.JWT([]byte(key)))
	e.DELETE("/users", handler.DeleteUser(), middleware.JWT([]byte(key)))
}

// registrasi add user
func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponses("berhasil register", ToResponse(res, "reg", "")))
	}

}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}
		cnv := ToDomain(input)
		log.Println("\n\n\ndata login \n", input, "\n\n")
		res, token, err := us.srv.Login(cnv.Email, cnv.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessLoginResponses("berhasil login", ToResponse(res, "login", token)))
	}
}

func (us *userHandler) EditProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponses(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		var input EditUserRequestFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}

		updateData := ToDomain(input)

		res, err := us.srv.UpdateProfile(updateData, c)

		log.Println("")
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponses("Success update data", ToResponse(res, "edit", "")))
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponses(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		if _, err := us.srv.DeleteProfile(c); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessDeleteResponses("Success Delete Data"))
	}
}

func (us *userHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponses(err.Error()))
		} else {
			log.Println("Authorized request.")
		}
		paramEmail := c.Param("email")
		log.Println("email awal:", paramEmail)
		// ID, err := strconv.Atoi(c.Param("email"))
		res, err := us.srv.Profile(paramEmail)
		log.Println("email", paramEmail)
		log.Println("res", res)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponses("sucses get userBy Email", ToResponse(res, "get", "")))
	}
}

func (us *userHandler) ShowAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllUser()
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponses("success get all user", ToResponse(res, "get", "")))
	}
}
