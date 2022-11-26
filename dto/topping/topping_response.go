package toppingdto

type ToppingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Qty   int    `json:"-" form:"qty"`
	Image string `json:"image"`
}

type ToppingUserResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Qty    int    `json:"-" form:"qty"`
	Image  string `json:"image"`
	UserID int    `json:"-"`
}

func (ToppingResponse) TableName() string {
	return "toppings"
}

func (ToppingUserResponse) TableName() string {
	return "toppings"
}
