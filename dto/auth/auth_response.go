package authdto

type RegisterResponse struct {
	Name string `gorm:"type: varchar(255)" json:"fullname"`
}

type LoginResponse struct {
	Id    int    `gorm:"type: int" json:"-"`
	Name  string `gorm:"type: varchar(255)" json:"fullname"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
