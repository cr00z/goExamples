package main

import (
	"github.com/julienschmidt/httprouter"

	"github.com/cr00z/goSite/mysqlSite/app/controller"
	"github.com/cr00z/goSite/mysqlSite/app/server"

	"log"
	"net/http"
)

func main() {
	err := server.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	r := httprouter.New()
	routes(r)
	if err := http.ListenAndServe("localhost:4444", r); err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
	r.POST("/user/add", controller.AddUser)
	r.DELETE("/user/delete/:userId", controller.DeleteUser)
	r.POST("/user/update/:userId", controller.UpdateUser)
}
