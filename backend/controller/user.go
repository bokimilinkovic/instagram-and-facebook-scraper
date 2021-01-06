package controller

import (
	"fmt"
	"holycode-task/model"
	"holycode-task/service"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// UserController represents controller for user apis.
type UserController struct {
	domain      string
	UserService *service.UserService
}

// NewUserController creates new UserController.
func NewUserController(domain string, us *service.UserService) *UserController {
	return &UserController{domain: domain, UserService: us}
}

// Create godoc.
// @Summary Login checks users credentials.
// @Description Login checks if the user is present in database, and returns encoded token with his informations.
// @Accept  json
// @Produce  json
// @Param RequestBody body dto.ProductDto true " "
// @Success 201 {object} string
// @Failure 400 "Bad_credentials_provided"
// @Router /v1/login [post]
func (uc *UserController) Login(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := uc.UserService.Authentificate(user.Username, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad credentials provided : %s", err)
	}

	//Creating token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//Encoding token and sending as response
	t, err := token.SignedString([]byte("some-secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}

// Create godoc.
// @Summary FindAll returns all users present in application.
// @Description FindAll returns all users.
// @Accept  json
// @Produce  json
// @Success 20 {array} []model.User
// @Failure 500 "Internal_server_error"
// @Router /v1/users [get]
func (uc *UserController) FindAll(c echo.Context) error {
	users, err := uc.UserService.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

// Create godoc.
// @Summary Register creates new user in database.
// @Description Register is storing new user which wants to use our application.
// @Accept  json
// @Produce  json
// @Param RequestBody body model.User true " "
// @Success 200 {object} string "Successfully_registered"
// @Failure 400,404 "Bad_credentials_provided"
// @Router /v1/register [post]
func (uc *UserController) Register(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(pass)
	err = uc.UserService.CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "successfully registered!")
}

func (uc *UserController) WhoAmI(c echo.Context) error {
	props, _ := c.Get("props").(jwt.MapClaims)
	// := c.Request().Context().Value("props").(jwt.MapClaims)
	ret := fmt.Sprintf("Hello %v : %v", props["username"], props["admin"])
	return c.String(http.StatusOK, ret)
}
