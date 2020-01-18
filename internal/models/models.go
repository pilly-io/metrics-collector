package models

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
}

type Namespace struct {
	gorm.Model
	Name      string
	Labels    string
}

type PodOwner struct {
	gorm.Model
	Type      string
	Name      string
	Labels    string
}
