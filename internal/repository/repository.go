package repository

import (
	"gorm.io/gorm"
)

// Repository is a structure that contains all repositories
type Repository struct {
	Warehouse *WarehouseRepository
}

// New creates a new Repository instance
func New(db *gorm.DB) *Repository {
	return &Repository{
		Warehouse: NewWarehouseRepository(db),
	}
}
