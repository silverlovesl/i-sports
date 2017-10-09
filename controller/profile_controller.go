package controller

import (
	"i-sports/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ProfileController struct {
	ProfileService *service.ProfileService
}

func NewProfileController() *ProfileController {
	return &ProfileController{
		ProfileService: service.NewProfileService(),
	}
}

func (pc *ProfileController) Login(c echo.Context) error {
	loginID := c.FormValue("loginId")
	password := c.FormValue("password")
	profile := pc.ProfileService.Login(loginID, password)
	return c.JSON(http.StatusOK, profile)
}

func (pc *ProfileController) Profile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid profile ID")
	}

	profile := pc.ProfileService.Profile(id)
	return c.JSON(http.StatusOK, profile)
}
