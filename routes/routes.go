package routes

import (
	"git.hifx.in/ad_engine/engine"
	"github.com/labstack/echo"
)

//Set sets all routes
func Set(e *echo.Echo) {

	//Get Token

	//Project Management
	e.POST("/recommend", engine.Recommend) //Listing all the release under a project
	// e.POST("/releases", handlers.AddReleases) //Adding all the release under a project

}
