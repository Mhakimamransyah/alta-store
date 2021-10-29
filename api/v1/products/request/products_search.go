package request

import (
	"altaStore/business/products"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewFilterProducts(c echo.Context) *products.FilterProducts {
	id_categories, err := strconv.Atoi(c.QueryParam("id_categories"))
	if err != nil {
		// set default 0 if categories not set
		id_categories = 0
	}
	query := c.QueryParam("query")

	sort := c.QueryParam("sort")

	price_max, err := strconv.Atoi(c.QueryParam("pmax"))
	if err != nil {
		// set default 0 if price max not set
		price_max = 0
	}

	price_min, err := strconv.Atoi(c.QueryParam("pmin"))
	if err != nil {
		// set default -99 if price min not set
		price_min = -1
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		// set default 0 if page not set
		page = 0
	}

	per_page, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		// set default 100 if page not set
		per_page = 100
	}

	return &products.FilterProducts{
		CategoriesId: id_categories,
		Query:        query,
		Sort:         sort,
		Price_max:    price_max,
		Price_min:    price_min,
		Page:         page,
		Per_page:     per_page,
	}
}
