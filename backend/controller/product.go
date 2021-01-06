package controller

import (
	"fmt"
	"holycode-task/controller/dto"
	"holycode-task/service"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// ProductController manages product requests.
type ProductController struct {
	product *service.ProductService
}

// NewProductController creates new product controller.
func NewProductController(pr *service.ProductService) *ProductController {
	return &ProductController{product: pr}
}

// Create godoc.
// @Summary Create creates new product.
// @Description Creates new product that is going to be saved in database.
// @Accept  json
// @Produce  json
// @Param RequestBody body dto.ProductDto true " "
// @Success 201 {object} model.Product
// @Failure 400,404 "Can_not_bind_request_Body"
// @Router /v1/products [post]
func (pc *ProductController) Create(c echo.Context) error {
	var product *dto.ProductDto
	// Source
	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Can not bind request body")
	}

	created, err := pc.product.CreateNew(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, created)
}

func (pc *ProductController) UploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./images/" + file.Filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, "Uploaded successfully")
}

// FindAllProducts godoc
// @Summary FindAll returns all products
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.Product
// @Failure 500 {object} string "Internal server error"
// @Router /v1/products [get]
func (pc *ProductController) FindAll(c echo.Context) error {
	products, err := pc.product.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("Error returning products : %v", err))
	}
	return c.JSON(http.StatusOK, products)
}
