package model

import "github.com/cr00z/goSite/mysqlSite/app/server"

type User struct {
	Id      int
	Name    string
	Surname string
}

func GetAllUsers() (users []User, err error) {
	query := "SELECT * FROM users"
	err = server.Db.Select(&users, query)
	return
}

func NewUser(name string, surname string) *User {
	return &User{Name: name, Surname: surname}
}

func GetUserById(userId string) (u User, err error) {
	query := "SELECT * FROM users WHERE id= ?"
	err = server.Db.Get(&u, query, userId)
	return
}

func (u *User) Add() (err error) {
	query := "INSERT INTO users (name, surname) VALUES (?, ?)"
	_, err = server.Db.Exec(query, u.Name, u.Surname)
	return
}

func (u *User) Delete() (err error) {
	query := "DELETE FROM users WHERE id = ?"
	_, err = server.Db.Exec(query, u.Id)
	return
}

func (u *User) Update() (err error) {
	query := "UPDATE users SET name = ?, surname = ? WHERE id = ?"
	_, err = server.Db.Exec(query, u.Name, u.Surname, u.Id)
	return
}
