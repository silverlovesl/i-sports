package main

import (
	"i-sports-backend/controller"
	"i-sports-backend/util"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

var engine *xorm.Engine

func main() {
	util.InitEngine()
	initRouter()
}

func initRouter() {
	e := echo.New()
	g := e.Group("/api/v1")
	ctrlProfile := controller.GetProfileController()
	ctrlMeasure := controller.GetMeasurementController()

	g.POST("/login", ctrlProfile.DoLogin)
	g.GET("/profile/:id", ctrlProfile.GetProfile)
	g.GET("/measure/:userId", ctrlMeasure.GetMeasurements)
	e.Logger.Fatal(e.Start(":3000"))
}

func test(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>API is working 😃</h1>")
}
