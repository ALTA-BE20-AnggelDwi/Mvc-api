package entities

type ProductCore struct {
	ID          uint         `json:"id" form:"id"`
	Name        string       `json:"name" form:"name"`
	UserID      uint         `json:"user_id" form:"user_id" gorm:"index"`
	Description string       `json:"description" form:"description"`
	User        UserResponse `gorm:"foreignKey:UserID"`
}

type Product struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id" gorm:"index"`
	Description string `json:"description" form:"description"`
}
