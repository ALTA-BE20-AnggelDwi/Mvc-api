package entities

import "time"

type UserCore struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	Password    string    `json:"password" form:"password"`
	Address     string    `json:"address" form:"address"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	Role        string    `json:"role" form:"role"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Role        string `json:"role" form:"role"`
}
