# Warehouse Management System

A simple warehouse management system built with Go, Gin, and PostgreSQL.

## Features

- Product management (CRUD operations)
- Transaction tracking (in/out)
- Statistics and reporting
- Admin-only access
- JWT authentication
- Swagger documentation

## Prerequisites

- Go 1.22 or higher
- PostgreSQL
- MinIO (optional, for file storage)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Avazbek-02/DE-Lider-Warehouse.git
cd DE-Lider-Warehouse
```

2. Install dependencies:
```bash
go mod download
```

3. Create and configure .env file:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run the application:
```bash
go run cmd/main.go
```

## API Documentation

The API documentation is available at `/swagger/index.html` when the server is running.

### Endpoints

#### Products
- `POST /api/products` - Create a new product
- `GET /api/products` - Get all products
- `GET /api/products/{id}` - Get a product by ID
- `PUT /api/products/{id}` - Update a product
- `DELETE /api/products/{id}` - Delete a product

#### Transactions
- `POST /api/transactions` - Create a new transaction

#### Statistics
- `GET /api/statistics` - Get warehouse statistics

## Authentication

All endpoints require admin authentication using JWT tokens. Include the token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## Database Schema

### Products
- ID (auto-generated)
- Name
- Description
- Price
- Quantity
- Unit
- Category
- CreatedAt
- UpdatedAt
- DeletedAt

### Transactions
- ID (auto-generated)
- ProductID (foreign key)
- Type (in/out)
- Quantity
- Date
- Description
- CreatedAt
- UpdatedAt
- DeletedAt

### Admins
- ID (auto-generated)
- Username
- Password
- CreatedAt
- UpdatedAt
- DeletedAt

## License

This project is licensed under the MIT License - see the LICENSE file for details.