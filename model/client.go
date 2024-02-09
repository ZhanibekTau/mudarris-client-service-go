package model

import "time"

type Client struct {
	Id        int        `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Lastname  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
	Token     string     `json:"token"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u *Client) TableName() string {
	return "clients"
}
