package main

import "github.com/vinshop/docer"

type UserProfile struct {
	Fullname string `json:"fullname"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type Pet struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type User struct {
	ID       uint         `json:"id"`
	Username string       `json:"username"`
	Profile  *UserProfile `json:"profile"`
	Pets     []*Pet       `json:"pets"`
}

func main() {
	doc := docer.Parse(User{})
	if err := doc.JSON("doc.json"); err != nil {
		panic(err)
	}
	if err := doc.Generate("doc.md"); err != nil {
		panic(err)
	}
}
