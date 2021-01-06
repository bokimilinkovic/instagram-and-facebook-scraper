package main

import (
	"fmt"
	"holycode-task/controller"
	"holycode-task/middleware"
	"holycode-task/repository/postgres"
	"holycode-task/service"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "holycode-task/docs"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"gopkg.in/tylerb/graceful.v1"
)

// @title Impact Bakery task API
// @version 1.0
// @description This is a server side of task.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env: %s", err)
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error converting port: %s", err)
	}

	config := postgres.PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	//Open database connection
	store, err := postgres.Open(config)
	if err != nil {
		log.Fatalf("Error openning connection to database: %s", err)
	}
	defer store.Close()
	if err = store.MigrateTables(); err != nil {
		log.Fatalf("Error migrating user table: %s", err)
	}

	// Setting up echo server
	e := echo.New()
	e.Use(echomiddleware.Logger(), echomiddleware.CORSWithConfig(
		echomiddleware.CORSConfig{
			AllowOrigins: []string{"*", "http://localhost:3000/"},
		},
	))

	userService := service.NewUserService(store)
	userController := controller.NewUserController("localhost", userService)
	userLoader := &middleware.UserLoader{}

	productService := service.NewProductService(store)
	productController := controller.NewProductController(productService)

	insta := goinsta.New(os.Getenv("INSTA_USERNAME"), os.Getenv("INSTA_PASSWORD"))
	err = insta.Login()
	if err != nil {
		fmt.Printf("Error occured during instagram login : %v", err)
	}
	defer insta.Logout()

	scraper := service.NewScraper(insta, store)
	scraperHadnler := controller.NewScraperHandler(scraper)

	//Serve static files: localhost:8080/static/image.png
	e.Static("/static", "images")

	e.GET("/v1/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/v1/users", userController.FindAll)
	e.POST("/v1/users/register", userController.Register)
	e.POST("/v1/users/login", userController.Login)
	e.GET("/v1/users/whoami", userController.WhoAmI, userLoader.Do)

	// Products
	e.POST("/v1/products", productController.Create, userLoader.Do)
	e.POST("/v1/image", productController.UploadImage)
	e.GET("/v1/products", productController.FindAll)

	// Instagram
	e.GET("/v1/instagram/:name", scraperHadnler.SearchInstagramByName)

	// Facebook
	e.GET("/v1/socialmedia/:name", scraperHadnler.SearchSocialMediaByName)

	e.Server.Addr = ":8080"
	graceful.DefaultLogger().Println("Application has successfully started at port: ", 8080)
	if err = graceful.ListenAndServe(e.Server, 5*time.Second); err != nil {
		log.Fatalf("Can not run server %s", err)
	}

}
