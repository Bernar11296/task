package main

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server is running")
	s := echo.New()

	s.GET("/status", Handler)
	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {

	d := time.Date(time.Now().Year(), time.September, 25, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	res := fmt.Sprintf("Количество дней до дня рождения: %d", int64(dur.Hours()/24))
	return ctx.String(200, res)
}
