package handler

import (
	"github.com/labstack/echo"
)

type SqlStruct struct{}

func (sql *SqlStruct) Index(c echo.Context) error {
	return c.String(200, "ok")
}

func sql2Struct(sql string) string {
	return ""
}
