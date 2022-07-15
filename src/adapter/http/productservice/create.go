package productservice

import (
	"encoding/json"
	"github.com/nelsonlpco/products-api/src/core/dto"
	"net/http"
)

func (s service) Create(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	productRequest, err := dto.FromJSONCreateProductRequest(request.Body)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	product, err := s.usecase.Create(ctx, productRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(product)
}
