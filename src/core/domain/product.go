package domain

import (
	"context"
	"github.com/nelsonlpco/products-api/src/core/dto"
	"net/http"
)

// Product is entity of table product database column
type Product struct {
	ID          int32   `json:"ID,omitempty"`
	Name        string  `json:"name,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
}

// ProductService is a contract of http adapter layer
type ProductService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type ProductUseCase interface {
	Create(ctx context.Context, productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(ctx context.Context, paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}

// ProductRepository is a contract of database connection adapter layer
type ProductRepository interface {
	Create(ctx context.Context, productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(ctx context.Context, paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}
