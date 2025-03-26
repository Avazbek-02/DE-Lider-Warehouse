package usecase

import (
	"errors"
	"time"

	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/entity"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// WarehouseUseCase implements business logic for warehouse operations
type WarehouseUseCase struct {
	repo   *repository.WarehouseRepository
	jwtKey []byte
	jwtTTL time.Duration
}

// NewWarehouseUseCase creates a new WarehouseUseCase instance
func NewWarehouseUseCase(
	repo *repository.WarehouseRepository,
	jwtSecret string,
	jwtTTL time.Duration,
) *WarehouseUseCase {
	return &WarehouseUseCase{
		repo:   repo,
		jwtKey: []byte(jwtSecret),
		jwtTTL: jwtTTL,
	}
}

// Product operations
func (uc *WarehouseUseCase) CreateProduct(product *entity.Product) error {
	return uc.repo.CreateProduct(product)
}

func (uc *WarehouseUseCase) GetProduct(id uint) (*entity.Product, error) {
	return uc.repo.GetProduct(id)
}

func (uc *WarehouseUseCase) UpdateProduct(product *entity.Product) error {
	return uc.repo.UpdateProduct(product)
}

func (uc *WarehouseUseCase) DeleteProduct(id uint) error {
	return uc.repo.DeleteProduct(id)
}

func (uc *WarehouseUseCase) GetAllProducts() ([]entity.Product, error) {
	return uc.repo.GetAllProducts()
}

// Transaction operations
func (uc *WarehouseUseCase) CreateTransaction(transaction *entity.Transaction) error {
	// Business logic validation
	if transaction.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	if transaction.Type != "in" && transaction.Type != "out" {
		return errors.New("transaction type must be 'in' or 'out'")
	}

	return uc.repo.CreateTransaction(transaction)
}

func (uc *WarehouseUseCase) GetTransactions(startDate, endDate time.Time) ([]entity.Transaction, error) {
	return uc.repo.GetTransactions(startDate, endDate)
}

// Statistics operations
func (uc *WarehouseUseCase) GetStatistics(startDate, endDate time.Time) (*entity.Statistics, error) {
	return uc.repo.GetStatistics(startDate, endDate)
}

// Admin authentication
func (uc *WarehouseUseCase) AdminLogin(username, password string) (string, error) {
	// Get admin from database
	admin, err := uc.repo.GetAdmin(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": admin.Username,
		"admin_id": admin.ID,
		"exp":      time.Now().Add(uc.jwtTTL).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString(uc.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
