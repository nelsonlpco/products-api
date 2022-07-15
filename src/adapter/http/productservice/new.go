package productservice

import "github.com/nelsonlpco/products-api/src/core/domain"

type service struct {
	usecase domain.ProductUseCase
}

func New(usecase domain.ProductUseCase) domain.ProductService {
	return &service{usecase: usecase}
}
