package entity

import (
	"github.com/google/uuid"
	"time"
)

type UserData struct {
	Id         uuid.UUID
	Username   string
	Name       string
	ChatId     int
	DeviceType DeviceType
	Token      string
	EndDate    time.Time
}

type DeviceType string

const (
	AndroidDevice DeviceType = "AndroidDevice"
	IOSDevice     DeviceType = "IOSDevice"
	PC            DeviceType = "PC"
)
