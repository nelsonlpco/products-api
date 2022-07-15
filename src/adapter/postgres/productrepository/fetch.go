package productrepository

import (
	"context"
	"github.com/nelsonlpco/products-api/src/core/domain"
	"github.com/nelsonlpco/products-api/src/core/dto"
	"strings"
)

func (r repository) Fetch(ctx context.Context, pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	products := []domain.Product{}
	total := int32(0)

	queryString := "SELECT * FROM product ORDER BY $1 OFFSET $2 LIMIT $3 "

	{
		rows, err := r.db.Query(
			ctx,
			queryString,
			strings.Join(pagination.Sort, ","),
			pagination.Page,
			pagination.ItemsPerPage,
		)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Product{}
			rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description)

			products = append(products, product)
		}
	}

	{
		err := r.db.QueryRow(ctx, "select COUNT(*) from product").Scan(&total)
		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.Product]{
		Items: products,
		Total: total,
	}, nil
}
