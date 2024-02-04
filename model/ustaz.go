package model

import "time"

type Ustaz struct {
	Id             int        `json:"id" gorm:"primary_key"`
	Name           string     `json:"name"`
	Lastname       string     `json:"lastname"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	Degree         string     `json:"degree"`
	University     string     `json:"university"`
	AdditionalInfo string     `json:"additional_info"`
	Experience     float32    `json:"experience"`
	Phone          string     `json:"phone"`
	Token          string     `json:"token"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

func (u *Ustaz) TableName() string {
	return "ustazs"
}
