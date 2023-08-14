package models

import (
	"database/sql"

	"github.com/todolist/pkg/config"
)

type App struct {
	DB     *sql.DB
	Config config.Config
}
