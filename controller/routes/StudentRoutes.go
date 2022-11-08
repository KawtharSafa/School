package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"school.com/packages/config"
)

func StudentRoutes(e *echo.Context, db *gorm.DB) {
	StudentController := config.WireStudentController(db)
	e.GET("/students", StudentController.GetAllStudents)
}
