package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	// Profile   ProfileResponse `json:"-"`
	// UserProduct []ProductUserResponse `json:"product"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
