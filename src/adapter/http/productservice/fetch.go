package productservice

import (
	"encoding/json"
	"github.com/nelsonlpco/products-api/src/core/dto"
	"net/http"
)

func (s service) Fetch(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	paginationRequest, err := dto.FromValuePaginationRequestParams(request)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	products, err := s.usecase.Fetch(ctx, paginationRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}
