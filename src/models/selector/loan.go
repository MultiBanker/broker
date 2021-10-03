package selector

import (
	"net/url"
	"strconv"
)

type Program struct {
	Amount    int
	Terms     []int
	SKU       string
	IsEnabled bool
}

func (p Program) HasAmount() bool {
	return p.Amount > 0
}

func (p Program) HasTerms() bool {
	return len(p.Terms) > 0
}

func (p Program) HasSKU() bool {
	return p.SKU != ""
}

const (
	pageKey  = "page"
	limitKey = "limit"
)

func ParsePaging(url url.Values) (page int64, limit int64, err error) {
	pageField := url.Get(pageKey)
	limitField := url.Get(limitKey)

	if pageField != "" {
		page, err = strconv.ParseInt(pageField, 10, 64)
		if err != nil {
			return
		}
	}
	if limitField != "" {
		limit, err = strconv.ParseInt(limitField, 10, 64)
		if err != nil {
			return
		}
	}

	return
}
