package handler

import (
	"github.com/labstack/echo"
)

type IndexStruct struct{}

func (i *IndexStruct) Index(c echo.Context) error {
	return c.String(200, "There is nothing")
}
