package models

import (
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Vendor      string `json:"vendor"`
	Name        string `json:"name"`
	Serial      string `json:"serial"`
	Version     string `json:"version"`
	Platform    string `json:"platform"`
	LastUpdated string `json:"LastUpdated"`
}
