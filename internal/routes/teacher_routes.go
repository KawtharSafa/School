package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"school.com/packages/config"
)

func TeacherRoutes(e *echo.Echo, db *gorm.DB) {
	TeacherController := config.WireTeacherController(db)
	e.GET("/teachers", TeacherController.GetAllTeachers)
}
