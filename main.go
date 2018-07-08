package main

import (
	"github.com/Laily123/tool-site/handler"
	"github.com/labstack/echo"
)

var (
	indexHandler *handler.IndexStruct
	sqlHandler   *handler.SqlStruct
)

func router(e *echo.Echo) {
	e.GET("/", indexHandler.Index)
	e.GET("/sql", sqlHandler.Index)
}

func main() {
	e := echo.New()
	router(e)
	e.Start(":8080")
}
