package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/letanthang/go_fw/db"
	"github.com/letanthang/go_fw/types"
)

func AddStudent(c echo.Context) error {
	var req types.AddStudentReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	result, err := db.AddStudent(req)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, result)
}
