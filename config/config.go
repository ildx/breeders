package config

import (
	"database/sql"
	"sync"

	"github.com/ildx/breeders/models"
)

type Application struct {
	Models *models.Models
}

var (
	db       *sql.DB
	instance *Application
	once     sync.Once
)

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models: models.New(db),
		}
	})
	return instance
}
