package models

import (
	usersdto "backend-WE/dto/users"
)

type Topping struct {
	ID     int                           `json:"id" gorm:"primary_key:auto_increment"`
	Title  string                        `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int                           `json:"price" form:"price" gorm:"type: int"`
	Qty    int                           `json:"-" form:"qty"`
	Image  string                        `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int                           `json:"-" form:"user_id"`
	User   usersdto.UsersProfileResponse `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
