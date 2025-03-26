package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null"`
	Quantity    int            `json:"quantity" gorm:"not null"`
	Unit        string         `json:"unit" gorm:"not null"`
	Category    string         `json:"category" gorm:"not null"`
}

type Transaction struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	ProductID   uint           `json:"product_id" gorm:"not null"`
	Product     Product        `json:"product" gorm:"foreignKey:ProductID"`
	Type        string         `json:"type" gorm:"not null"` // "in" or "out"
	Quantity    int            `json:"quantity" gorm:"not null"`
	Date        time.Time      `json:"date" gorm:"not null"`
	Description string         `json:"description"`
}

type Statistics struct {
	TotalProducts    int       `json:"total_products"`
	TotalValue       float64   `json:"total_value"`
	MonthlyIn        int       `json:"monthly_in"`
	MonthlyOut       int       `json:"monthly_out"`
	MonthlyValueIn   float64   `json:"monthly_value_in"`
	MonthlyValueOut  float64   `json:"monthly_value_out"`
	LowStockProducts []Product `json:"low_stock_products"`
}

type Admin struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
}
