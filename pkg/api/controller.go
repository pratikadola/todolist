package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/todolist/pkg/models"
	"github.com/todolist/pkg/service"
)

type TodolistController struct {
	service *service.Service
}

func NewTodolistController(service *service.Service) *TodolistController {
	return &TodolistController{service: service}
}

func (c *TodolistController) CreateTodoList(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	body, readErr := ioutil.ReadAll(request.Body)
	if readErr != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	} else {
		list := models.List{}
		jsonErr := json.Unmarshal(body, &list)
		if jsonErr != nil {
			log.Fatalln("There was an error encoding the initialized struct", jsonErr)
		}
		resp, err := c.service.CreateNewList(list)
		if err != nil {
			log.Fatalln("Error creating list", err)
		}
		writer.Write([]byte(string(resp.ID)))
	}
}

func (c *TodolistController) AddTaskToList(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	body, readErr := ioutil.ReadAll(request.Body)
	if readErr != nil {
		log.Fatalln("There was an error decoding the request body into the struct", readErr)
	} else {
		task := models.Task{}
		jsonErr := json.Unmarshal(body, &task)
		if jsonErr != nil {
			log.Fatalln("There was an error encoding the initialized struct", jsonErr)
		}
		listID := mux.Vars(request)["listid"]
		lID, err := strconv.Atoi(listID)
		if err != nil {
			log.Fatalln("error fetching listID", err)
		}
		resp, err := c.service.AddTaskToList(task, lID)
		if err != nil {
			log.Fatalln("Error creating list")
		}
		writer.Write([]byte(string(resp.ID)))
	}
}
