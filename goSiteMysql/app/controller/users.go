package controller

import (
	"encoding/json"

	"github.com/julienschmidt/httprouter"

	"github.com/cr00z/goSite/mysqlSite/app/model"

	"html/template"
	"net/http"
	"path/filepath"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	main := filepath.Join("public", "html", "usersDynamicPage.html")
	common := filepath.Join("public", "html", "common.html")
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	if name == "" || surname == "" {
		http.Error(rw, "Name or surname is empty", 400)
		return
	}
	user := model.NewUser(name, surname)
	err := user.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = json.NewEncoder(rw).Encode("User added")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = user.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = json.NewEncoder(rw).Encode("User deleted")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	user.Name = name
	user.Surname = surname

	err = user.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = json.NewEncoder(rw).Encode("User updated")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
