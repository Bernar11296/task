package endpoint

import (
	"fmt"
	"text/template"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft(days, month string) int64
}
type Endpoint struct {
	s Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{
		s: svc,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	c := ctx.Request()
	dat := c.FormValue("day")
	month := c.FormValue("month")

	d := e.s.DaysLeft(dat, month)
	res := fmt.Sprintf("Количество дней до дня рождения: %d", d)
	return ctx.String(200, res)
}

func (e *Endpoint) Hello(ctx echo.Context) error {
	tmpl, _ := template.ParseFiles("ui/static/index.html")
	return tmpl.Execute(ctx.Response().Writer, nil)

}
