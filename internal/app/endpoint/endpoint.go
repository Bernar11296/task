package endpoint

import (
	"fmt"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int64
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
	d := e.s.DaysLeft()
	res := fmt.Sprintf("Количество дней до дня рождения: %d", d)
	return ctx.String(200, res)
}
