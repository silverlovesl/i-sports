package controller

import (
	"i-sports/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// MeasurementController [TODO:Comment]
type MeasurementController struct {
	MeasurementService *service.MeasurementService
}

// NewMeasurementController [TODO:Comment]
func NewMeasurementController() *MeasurementController {
	return &MeasurementController{
		MeasurementService: service.NewMeasurementService(),
	}
}

// Measurement [TODO:Comment]
func (mc *MeasurementController) Measurement(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid userId")
	}
	measurements := mc.MeasurementService.GetMeasurements(userID)
	return c.JSON(http.StatusOK, WrapArray(measurements))
}
