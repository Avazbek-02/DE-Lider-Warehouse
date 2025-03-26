package repository

import (
	"errors"
	"time"

	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/entity"
	"gorm.io/gorm"
)

// Custom errors
var (
	ErrInsufficientStock = errors.New("insufficient stock quantity")
)

type WarehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

// Product operations
func (r *WarehouseRepository) CreateProduct(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *WarehouseRepository) GetProduct(id uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *WarehouseRepository) UpdateProduct(product *entity.Product) error {
	// Only update non-empty fields
	updates := make(map[string]interface{})
	if product.Name != "" {
		updates["name"] = product.Name
	}
	if product.Description != "" {
		updates["description"] = product.Description
	}
	if product.Price != 0 {
		updates["price"] = product.Price
	}
	if product.Quantity != 0 {
		updates["quantity"] = product.Quantity
	}
	if product.Unit != "" {
		updates["unit"] = product.Unit
	}
	if product.Category != "" {
		updates["category"] = product.Category
	}
	return r.db.Model(product).Updates(updates).Error
}

func (r *WarehouseRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}

func (r *WarehouseRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, err
}

// Transaction operations
func (r *WarehouseRepository) CreateTransaction(transaction *entity.Transaction) error {
	// Start transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Update product quantity
	var product entity.Product
	if err := tx.First(&product, transaction.ProductID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if transaction.Type == "in" {
		product.Quantity += transaction.Quantity
	} else if transaction.Type == "out" {
		if product.Quantity < transaction.Quantity {
			tx.Rollback()
			return ErrInsufficientStock
		}
		product.Quantity -= transaction.Quantity
	}

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create transaction record
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *WarehouseRepository) GetTransactions(startDate, endDate time.Time) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&transactions).Error
	return transactions, err
}

// Statistics operations
func (r *WarehouseRepository) GetStatistics(startDate, endDate time.Time) (*entity.Statistics, error) {
	var stats entity.Statistics

	// Get total products and value
	var products []entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	stats.TotalProducts = len(products)
	for _, p := range products {
		stats.TotalValue += p.Price * float64(p.Quantity)
	}

	// Get monthly transactions
	var transactions []entity.Transaction
	if err := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&transactions).Error; err != nil {
		return nil, err
	}

	for _, t := range transactions {
		if t.Type == "in" {
			stats.MonthlyIn += t.Quantity
			stats.MonthlyValueIn += t.Product.Price * float64(t.Quantity)
		} else {
			stats.MonthlyOut += t.Quantity
			stats.MonthlyValueOut += t.Product.Price * float64(t.Quantity)
		}
	}

	// Get low stock products (less than 10)
	if err := r.db.Where("quantity < ?", 10).Find(&stats.LowStockProducts).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

// Admin operations
func (r *WarehouseRepository) GetAdmin(username string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return &admin, err
}
