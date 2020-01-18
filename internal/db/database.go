package main

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func NewDb(DbUri string) *gorm.DB {
	db, err := gorm.Open("sqlite3", DbUri)
	if err != nil {
		log.Error("Failed to connect database")
		os.Exit(1)
	}
	defer db.Close()
	// Temporary migration 
	//db.AutoMigrate(&Node{}, &Namespace{}, &PodOwner{})
	return db
}
