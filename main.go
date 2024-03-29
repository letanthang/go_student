package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/letanthang/go_student/config"
	"github.com/letanthang/go_student/db"
	"github.com/letanthang/go_student/route"
)

func main() {
	fmt.Println(config.Config)
	fmt.Println(db.Client)
	e := echo.New()
	route.All(e)

	err := e.Start(":9090")

	if err != nil {
		fmt.Println("Cannot start server", err)
	}

}
