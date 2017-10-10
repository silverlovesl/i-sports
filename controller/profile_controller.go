package controller

import (
	"i-sports/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ProfileController [TODO:Comment]
type ProfileController struct {
	ProfileService *service.ProfileService
}

// NewProfileController [TODO:Comment]
func NewProfileController() *ProfileController {
	return &ProfileController{
		ProfileService: service.NewProfileService(),
	}
}

// DoLogin [TODO:Comment]
func (pc *ProfileController) DoLogin(c echo.Context) error {
	loginID := c.FormValue("loginId")
	password := c.FormValue("password")
	profile := pc.ProfileService.DoLogin(loginID, password)
	return c.JSON(http.StatusOK, profile)
}

// GetProfile [TODO:Comment]
func (pc *ProfileController) GetProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid profile ID")
	}
	profile := pc.ProfileService.GetProfile(id)
	return c.JSON(http.StatusOK, profile)
}
