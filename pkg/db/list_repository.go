package db

import "github.com/todolist/pkg/models"

func (r *Repository) CreateNewList(l *models.List) (*models.List, error) {
	result := r.db.Create(l)
	if result.Error != nil {
		return nil, result.Error
	}
	return l, nil
}

func (r *Repository) AddTaskToList(task *models.Task) (*models.Task, error) {
	result := r.db.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}
