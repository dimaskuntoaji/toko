package entity

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserName  string    `json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterInput struct {
	UserName string `json:"username" binding:"required" example:"dimas"`
	Email    string `json:"email" binding:"required,email" example:"kunto@aji.com"`
	Password string `json:"password" binding:"required,min=6" example:"dimask"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email" example:"kunto@aji.com"`
	Password string `json:"password" binding:"required,min=6" example:"dimask"`
}

type UserResponseRegister struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}


type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
