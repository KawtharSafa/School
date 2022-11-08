package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"school.com/packages/config"
)

func ClassTeacherRoutes(e *echo.Context, db *gorm.DB) {
	ClassTeacherController := config.WireClassTeacherController(db)
	e.GET("/classes", ClassTeacherController.GetClassTeacher)
}
