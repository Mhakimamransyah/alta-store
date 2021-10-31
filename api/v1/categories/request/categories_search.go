package request

import (
	"altaStore/business/categories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewFilterProducts(c echo.Context) *categories.FilterCategories {
	query := c.QueryParam("query")

	id_admin, err := strconv.Atoi(c.QueryParam("id_admin"))
	if err != nil {
		id_admin = 0
	}
	status := c.QueryParam("status")
	if status == "" {
		status = "active"
	}
	sortName := c.QueryParam("sort_name")

	sortDate := c.QueryParam("sort_date")

	limit, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		offset = 0
	}

	return &categories.FilterCategories{
		Query:    query,
		AdminId:  id_admin,
		Status:   status,
		SortName: sortName,
		SortDate: sortDate,
		Offset:   offset,
		Limit:    limit,
	}
}
