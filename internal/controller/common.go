package controller

import "time"

type DTO struct {
	Name       string
	Username   string
	ChatId     int
	EndDate    time.Time
	DeviceType string
	Token      string
}
