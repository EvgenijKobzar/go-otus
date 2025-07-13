package model

import "time"

type Account struct {
	Id        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Evgenij"`
	FirstName string    `json:"first_name" example:"Kobzar"`
	LastName  string    `json:"last_name" example:""`
	Login     string    `json:"login" example:"ekobzar"`
	Password  string    `json:"password" example:"123456"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) GetId() int {
	return a.Id
}

func (a *Account) SetId(id int) {
	a.Id = id
}
