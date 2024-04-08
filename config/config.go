package config

import (
	"database/sql"
	"sync"

	"github.com/ildx/breeders/adapters"
	"github.com/ildx/breeders/models"
)

type Application struct {
	Models     *models.Models
	CatService *adapters.RemoteService
}

var (
	db         *sql.DB
	catService *adapters.RemoteService
	instance   *Application
	once       sync.Once
)

func New(pool *sql.DB, a *adapters.RemoteService) *Application {
	db = pool
	catService = a
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			CatService: catService,
			Models:     models.New(db),
		}
	})
	return instance
}
