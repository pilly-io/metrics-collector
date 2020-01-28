package models

import (
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
	UID          string
	Version      string
	OS           string
	Labels       string
}

type Namespace struct {
	gorm.Model
	Name   string
	Labels string
}

type Pod struct {
	gorm.Model
	Name      string
	Namespace string
	Labels    string
	OwnerType string
	OwnerName string
}

type PodMetric struct {
	gorm.Model
	MetricName  string
	MetricValue float64
	Namespace   string
	NodeName    string
	PodName     string
	OwnerType   string
	OwnerName   string
}
