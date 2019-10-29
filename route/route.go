package route

import (
	"github.com/labstack/echo"
	"github.com/letanthang/go_fw/handler"
)

func All(e *echo.Echo) {
	Public(e)
	Staff(e)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/student/v1/public")
	g.GET("/health", handler.CheckHealth)
	g.GET("/student", handler.GetAllStudent)

}

func Staff(e *echo.Echo) {
	g := e.Group("/api/student/v1/staff")
	g.POST("/student", handler.AddStudent)
}

func Private() {

}
