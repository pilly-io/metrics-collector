package main

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Node struct {
	gorm.Model
	Name         string
	InstanceType string
	Region       string
	Zone         string
	Hostname     string
	uid          uuid.UUID `gorm:"type:uuid"`
	Version      string
	OS           string
	Labels       string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type Namespace struct {
	gorm.Model
	Name      string
	Labels    string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type PodOwner struct {
	gorm.Model
	Type      string
	Name      string
	Labels    string
	UpdatedAt time.Time
	CreatedAt time.Time
}
