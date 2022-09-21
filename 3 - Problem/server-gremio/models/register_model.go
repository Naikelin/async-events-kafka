package models

type PatentToRegister struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Patent string `json:"patent"`
}
