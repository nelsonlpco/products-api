package di

import (
	"github.com/nelsonlpco/products-api/src/adapter/http/productservice"
	"github.com/nelsonlpco/products-api/src/adapter/postgres"
	"github.com/nelsonlpco/products-api/src/adapter/postgres/productrepository"
	"github.com/nelsonlpco/products-api/src/core/domain"
	"github.com/nelsonlpco/products-api/src/core/domain/productusecase"
)

func ConfigureProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
