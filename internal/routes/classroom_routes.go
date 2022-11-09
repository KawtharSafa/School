package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"school.com/packages/config"
)

func ClassroomRoutes(e *echo.Echo, db *gorm.DB) {
	ClassroomController := config.WireClassroomController(db)
	e.GET("/classes", ClassroomController.GetClassrooms)
}
