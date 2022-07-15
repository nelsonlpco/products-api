package productrepository

import (
	"context"
	"github.com/nelsonlpco/products-api/src/core/domain"
	"github.com/nelsonlpco/products-api/src/core/dto"
)

func (r repository) Create(ctx context.Context, productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product := domain.Product{}

	query := "INSERT INTO product (name, price, description) values ($1, $2, $3) returning *"
	err := r.db.QueryRow(
		ctx,
		query,
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
