package db

import (
	"fmt"

	"github.com/todolist/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	config config.Config
}

func NewRepository(conf config.Config) *Repository {
	db := connect(conf)
	return &Repository{db: db, config: conf}
}

func connect(conf config.Config) *gorm.DB {
	username := conf.Database.Username
	password := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port
	dbName := conf.Database.DBName
	pqConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)

	db, err := gorm.Open(postgres.Open(pqConnection))
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(2)

	return db
}

func (r *Repository) DropTableIfExists(name interface{}) error {
	return r.db.Migrator().DropTable(name)
}

func (r *Repository) Automigrate(name interface{}) error {
	return r.db.AutoMigrate(name)
}
