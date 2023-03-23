package service

import (
	"strconv"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft(days, month string) int64 {
	mt, _ := strconv.Atoi(month)
	day, _ := strconv.Atoi(days)
	var m = time.Month(mt)
	d := time.Date(time.Now().Year(), m, day, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	return int64(dur.Hours()) / 24
}
