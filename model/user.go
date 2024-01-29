package model

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "users"
}
