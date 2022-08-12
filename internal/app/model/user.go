package model

type User struct {
	Id     int
	Name   string
	Email  string
	salt   string
	digest string
}
