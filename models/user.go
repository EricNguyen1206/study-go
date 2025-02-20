package models

type UserModel struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Profile  string `json:"profile"`
}

type UserLoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignUpDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Profile  string `json:"profile"`
}

type UserUpdateProfileDto struct {
	ID      int    `json:"id" binding:"required"`
	Profile string `json:"profile" binding:"max=6"`
}
