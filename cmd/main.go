package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/todolist/pkg/api"
	"github.com/todolist/pkg/config"
	"github.com/todolist/pkg/db"
	"github.com/todolist/pkg/migration"
	"github.com/todolist/pkg/service"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	repo := db.NewRepository(cfg)
	migration.CreateDatabase(cfg, *repo)
	service := service.NewService(repo)
	controller := api.NewTodolistController(service)
	Init(controller)
}

func Init(controller *api.TodolistController) {
	r := mux.NewRouter()
	r.HandleFunc("/list", controller.CreateTodoList).Methods("POST")
	r.HandleFunc("/list/add-task/{listid}", controller.AddTaskToList).Methods("POST")
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
