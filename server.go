package main

import (
	"fmt"
	"log"
	"spamtrawler/app"
	"spamtrawler/app/routes"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Will print: "[date] [time] [filename]:[line]: [text]"
	//log.Println("Logging w/ line numbers on golangcode.com")

	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fmt.Println(app.MongoDB)

	routes.RouteHandler(e)
}
