package models

import "time"

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	// Profile   ProfileResponse `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
