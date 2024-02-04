package utils

import (
	"net/http"
	"strconv"
	"user-service-go/model"
)

// GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(r *http.Request) model.Pagination {
	// Initializing default
	//	var mode string
	limit := 30
	page := 1
	sort := "created_at asc"
	query := r.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return model.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
