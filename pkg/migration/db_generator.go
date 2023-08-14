package migration

import (
	"github.com/todolist/pkg/config"
	"github.com/todolist/pkg/db"
	"github.com/todolist/pkg/models"
)

func CreateDatabase(cfg config.Config, repo db.Repository) {
	if cfg.Database.Migration {
		// err := repo.DropTableIfExists(&models.List{})
		// if err != nil {
		// 	panic(err)
		// }
		// err = repo.DropTableIfExists(&models.Task{})
		// if err != nil {
		// 	panic(err)
		// }
		err := repo.Automigrate(&models.Task{})
		if err != nil {
			panic(err)
		}
		err = repo.Automigrate(&models.List{})
		if err != nil {
			panic(err)
		}

	}
}
