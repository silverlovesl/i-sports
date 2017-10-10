package controller

import (
	"i-sports/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// MeasurementController [TODO:Comment]
type measurementController struct{}

var measurementControllerIns = new(measurementController)

// GetMeasurementController [TODO:Comment]
func GetMeasurementController() *measurementController {
	return measurementControllerIns
}

// GetMeasurements [TODO:Comment]
func (mc *measurementController) GetMeasurements(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid userId")
	}
	measurements := service.GetMeasurementService().GetMeasurements(userID)
	return c.JSON(http.StatusOK, WrapArray(measurements))
}
