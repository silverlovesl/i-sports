package controller

import (
	"fmt"
	"i-sports/service"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// ProfileController [TODO:Comment]
type profileController struct {
}

var profileControllerIns = new(profileController)

// NewProfileController [TODO:Comment]
func GetProfileController() *profileController {
	if profileControllerIns == nil {
		time.Sleep(1 * time.Second)
		fmt.Println("A new profileController has been create ")
	}
	return profileControllerIns
}

// DoLogin [TODO:Comment]
func (pc *profileController) DoLogin(c echo.Context) error {
	loginID := c.FormValue("loginId")
	password := c.FormValue("password")
	profile := service.GetProfileServiceIns().DoLogin(loginID, password)
	return c.JSON(http.StatusOK, profile)
}

// GetProfile [TODO:Comment]
func (pc *profileController) GetProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid profile ID")
	}
	profile := service.GetProfileServiceIns().GetProfile(id)
	return c.JSON(http.StatusOK, profile)
}
