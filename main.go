package main

import (
	"github.com/ad_engine/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//main function
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))
	routes.Set(e)

	// engine.Recommend(e)
	// err = services.Load()
	// if err != nil {
	// 	logger.Error.Error("Error while loading services")
	// 	return
	// }

	e.Logger.Fatal(e.Start(":8080"))
}
