package routes

import (
	"github/ad_engine/engine"
	"github.com/labstack/echo"
)

//Set sets all routes
func Set(e *echo.Echo) {

	// Code to get the token

	//Project Management
	e.POST("/recommend", engine.Recommend) //Listing all the release under a project
	e.POST("/samplePrint", engine.Sample)
}
