package controller

import (
	"holycode-task/controller/dto"
	"holycode-task/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ScraperHandler represents handler that manages instagram and facebook scraping.
type ScraperHandler struct {
	Scraper *service.Scraper
}

// NewScraperHandler creates new ScraperHandler.
func NewScraperHandler(s *service.Scraper) *ScraperHandler {
	return &ScraperHandler{Scraper: s}
}

// Create godoc.
// @Summary SearchInstagramByName returns instagram datas based on provided name.
// @Accept  json
// @Produce  json
// @Param name path string true "Username in instagram"
// @Success 200 {object} model.InstagramAccount
// @Router /v1/instagram/:name [post]
func (sh *ScraperHandler) SearchInstagramByName(c echo.Context) error {
	name := c.Param("name")
	instaResponse, err := sh.Scraper.SearchInstagramByUsername(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, instaResponse)
}

// Create godoc.
// @Summary SearchSocialMediaByName returns instagram and facebook datas based on provided name.
// @Accept  json
// @Produce  json
// @Param name path string true "Username_in_instagram."
// @Success 200 {object} dto.SocialMediaDto
// @Failure 400 "Bad_request_payload"
// @Router /v1/socialmedia/:name [post]
func (sh *ScraperHandler) SearchSocialMediaByName(c echo.Context) error {
	name := c.Param("name")
	instaResponse, err := sh.Scraper.SearchInstagramByUsername(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	fbResponse, err := sh.Scraper.ScrapeFacebookProfile(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	socialMediaResponse := dto.CreateSocialMediaDto(instaResponse, fbResponse)

	return c.JSON(http.StatusOK, socialMediaResponse)
}
