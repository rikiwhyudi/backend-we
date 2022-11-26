package usersdto

type UserResponse struct {
	ID int `json:"id"`
	// Name string `json:"name" form:"name" validate:"required"`
	// // Email    string `json:"email" form:"email" validate:"required"`
	// // Password string `json:"password" form:"password" validate:"required"`
	// Token string `gorm:"type: varchar(255)" json:"token"`
}

type UsersProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"fullname"`
	Email string `json:"email"`
	Image string `json:"image"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
