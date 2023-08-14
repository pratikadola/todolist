package service

import (
	"github.com/todolist/pkg/db"
	"github.com/todolist/pkg/models"
)

type Service struct {
	repository *db.Repository
}

func NewService(repo *db.Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) CreateNewList(list models.List) (*models.List, error) {
	return s.repository.CreateNewList(&list)
}

func (s *Service) AddTaskToList(task models.Task, listID int) (*models.Task, error) {
	task.ListRefer = uint(listID)
	return s.repository.AddTaskToList(&task)
}
