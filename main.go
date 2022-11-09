package main

import (
	"fmt"
	"net/http"
	"school.com/packages/config"
	"school.com/packages/controller"
	controller2 "school.com/packages/internal/controller"
	"school.com/packages/internal/interface"

	// "school.com/packages/wire"

	// "github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func InitializeEvent() Event {
// 	wire.Build(GetMessage, GetGreeter, GetEvent)
// 	return Event{}
//  }

func main() {

	welcoming := fmt.Sprintf("Hello, %s", "Kawthar")
	fmt.Println(welcoming)

	// Echo instance
	e := echo.New()
	config.NewDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", heyAgain)

	e.GET("/students", controller.Students)
	e.GET("/teachers", controller.Teachers)
	e.GET("/classes", controller2.Classrooms)
	e.GET("/count", _interface.Count)

	// Start server
	e.Logger.Fatal(e.Start(":5500"))

	fmt.Println("connected to db:)")

	//for wire
	// event := InitializeEvent()
	// event.Start()

}

// Handle request by welcome msg
func heyAgain(c echo.Context) error {
	return c.String(http.StatusOK, "Hey Again From Here")
}
