package service_model

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (p Pagination) Parse(r *http.Request) (Pagination, error) {
	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return p, nil
		}

		p.Limit = l
	}

	offset := qs.Get("offset")

	if offset != "" {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return p, nil
		}

		p.Offset = l
	}

	sort := qs.Get("sort")
	if sort != "" {
		p.Sort = sort
	}

	return p, nil
}
