package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pilly-io/metrics-collector/internal/models"
)

// Database is wrapper for the orm
type Database struct {
	client *gorm.DB
}

// New creates an new DB object
func New(DBURI string) (Database, error) {
	db, err := gorm.Open("sqlite3", DBURI)
	return Database{client: db}, err
}

// Migrate sync the schemas of the DB
func (db *Database) Migrate() {
	// Dummy migration
	db.client.AutoMigrate(&models.Node{}, &models.Namespace{}, &models.Pod{}, &models.PodMetric{})
}

// Insert creates a new record in the right table
func (db *Database) Insert(value interface{}) {
	db.client.Create(value)
}
