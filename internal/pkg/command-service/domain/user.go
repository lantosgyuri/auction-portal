package domain

import "gorm.io/gorm"

type UserEvent interface {
	GetName() string
}

type CreateUserRequested struct {
	Name     string
	Password string
}

type DeleteUserRequest struct {
	Id   int
	Name string
}

type User struct {
	gorm.Model
	Id       int `gorm:"primaryKey"`
	Name     string
	Password string
}

type UserEventRaw struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	EventType string
	UserId    int
	Name      string
	Password  string
}

func (c CreateUserRequested) GetName() string {
	return c.Name
}

func (d DeleteUserRequest) GetName() string {
	return d.Name
}
