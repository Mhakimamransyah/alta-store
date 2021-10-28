package categories

import (
	"altaStore/api/common"
	"altaStore/api/middleware"
	"altaStore/business/categories"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	categories_service categories.Service
}

func InitCategoriesController(service categories.Service) *Controller {
	return &Controller{
		categories_service: service,
	}
}

func (service *Controller) GetAllCategoriesController(c echo.Context) error {
	page, errorPage := strconv.Atoi(c.QueryParam("page"))
	per_page, errorPerPage := strconv.Atoi(c.QueryParam("per_page"))
	if errorPage != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Page query params undefined"))
	}
	if errorPerPage != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Per Page query params undefined"))
	}
	data, err := service.categories_service.FindAllCategories(per_page, page)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}

func (service *Controller) GetAllSubCategoriesController(c echo.Context) error {
	page, errorPage := strconv.Atoi(c.QueryParam("page"))
	per_page, errorPerPage := strconv.Atoi(c.QueryParam("per_page"))
	if errorPage != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Page query params undefined"))
	}
	if errorPerPage != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Per Page query params undefined"))
	}
	id_categories, err := strconv.Atoi(c.Param("id_categories"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	data, err := service.categories_service.FindAllSubCategories(id_categories, per_page, page)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}

func (service *Controller) InsertCategoriesController(c echo.Context) error {
	categories := categories.CategoriesSpec{}
	c.Bind(&categories)

	username_admin := middleware.ExtractTokenKey(c, "username").(string)

	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	err = service.categories_service.InsertCategories(
		categories,
		id_admin,
		username_admin)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (service *Controller) ModifyCategoriesController(c echo.Context) error {
	categories := categories.CategoriesUpdatable{}
	c.Bind(&categories)
	id_categories, err := strconv.Atoi(c.Param("id_categories"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	username_admin := middleware.ExtractTokenKey(c, "username").(string)

	err = service.categories_service.ModifyCategories(
		categories,
		id_categories,
		id_admin,
		username_admin,
	)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (service *Controller) RemoveCategoriesController(c echo.Context) error {
	id_categories, err := strconv.Atoi(c.Param("id_categories"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	username_admin := middleware.ExtractTokenKey(c, "username").(string)

	err = service.categories_service.RemoveCategories(
		id_categories,
		id_admin,
		username_admin)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}
