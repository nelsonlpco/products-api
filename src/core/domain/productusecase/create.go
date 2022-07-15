package productusecase

import (
	"context"
	"github.com/nelsonlpco/products-api/src/core/domain"
	"github.com/nelsonlpco/products-api/src/core/dto"
)

func (u usecase) Create(ctx context.Context, productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := u.repository.Create(ctx, productRequest)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u usecase) Fetch(ctx context.Context, paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	products, err := u.repository.Fetch(ctx, paginationRequest)
	if err != nil {
		return nil, err
	}

	return products, nil
}
