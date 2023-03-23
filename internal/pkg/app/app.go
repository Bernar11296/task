package app

import (
	"fmt"
	"log"

	"github.com/Bernar11296/task/internal/app/endpoint"
	"github.com/Bernar11296/task/internal/app/mw"
	"github.com/Bernar11296/task/internal/app/service"
	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() *App {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()
	a.echo.GET("/hello", a.e.Hello)
	a.echo.GET("/status", a.e.Status, mw.RoleCheck)
	a.echo.POST("/status", a.e.Status, mw.RoleCheck)
	return a
}

func (a *App) Run() error {
	fmt.Println("server is running")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
