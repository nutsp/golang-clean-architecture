package database

import (
	"github.com/nutsp/golang-clean-architecture/config"
	"github.com/nutsp/golang-clean-architecture/pkg/datasource"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type IDatabase interface {
	OllamaDB() datasource.DB
	GptDB() datasource.DB
}

type Database struct {
	ollamadb *gorm.DB
	gptdb    *gorm.DB
}

func NewDatabase(cfg *config.Config) *Database {
	return &Database{
		ollamadb: lo.Must(datasource.NewDatabase(cfg.Database)),
		gptdb:    lo.Must(datasource.NewDatabase(cfg.Database)),
	}
}

func (db *Database) OllamaDB() datasource.DB {
	return &datasource.GormDB{db.ollamadb}
}

func (db *Database) GptDB() datasource.DB {
	return &datasource.GormDB{db.gptdb}
}
