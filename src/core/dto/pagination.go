package dto

import (
	"net/http"
	"strconv"
	"strings"
)

// PaginationRequestParams is an representation query string params
type PaginationRequestParams struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Sort         []string `json:"sort"`
}

// FromValuePaginationRequestParams converts query string params to a PaginationRequestParams struct
func FromValuePaginationRequestParams(request *http.Request) (*PaginationRequestParams, error) {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))

	paginationRequestParams := PaginationRequestParams{
		Search:       request.FormValue("search"),
		Descending:   strings.Split(request.FormValue("descending"), ","),
		Page:         page,
		ItemsPerPage: itemsPerPage,
		Sort:         strings.Split(request.FormValue("sort"), ","),
	}

	return &paginationRequestParams, nil
}
