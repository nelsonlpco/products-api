package productrepository

import (
	"github.com/nelsonlpco/products-api/src/adapter/postgres"
	"github.com/nelsonlpco/products-api/src/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
