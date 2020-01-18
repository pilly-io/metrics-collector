package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pilly-io/metrics-collector/internal/models"
)

type Database struct{
	client *gorm.DB
}

func NewDb(DBURI string) (Database, error) {
	db, err := gorm.Open("sqlite3", DBURI)
	return Database{client: db}, err
}

func (db Database) Migrate() {
	// Dummy migration
	db.client.AutoMigrate(&models.Node{}, &models.Namespace{}, &models.PodOwner{})
}