package models

import "gorm.io/gorm"

type Interface struct {
	gorm.Model
	Name        string   `json:"name"`
	MAC         string   `json:"mac"`
	Description string   `json:"description"`
	Driver      string   `json:"driver"`
	ARP         ARP      `json:"arp"`
	IPInfo      []IPInfo `json:"ipinfo"`
	Parent      string   `json:"parent"`
}

type IPInfo struct {
	gorm.Model
	InterfaceID uint
	Address     string `json:"address"`
	Bits        string `json:"bits"`
	Network     string `json:"network"`
}

type ARP struct {
	gorm.Model
	InterfaceID uint
}
