package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	// ID          int
	Username    string `json:"username" form:"username"`
	Displayname string `json:"displayname" form:"displayname"`
	Location    string `json:"location" form:"location"`
}
