package categories

import (
	"altaStore/api/common"
	"altaStore/api/middleware"
	"altaStore/api/v1/categories/request"
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
	filter := request.NewFilterProducts(c)
	data, err := service.categories_service.FindAllCategories(filter)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}

func (service *Controller) GetAllSubCategoriesController(c echo.Context) error {
	filter := request.NewFilterProducts(c)
	id_categories, err := strconv.Atoi(c.Param("id_categories"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	data, err := service.categories_service.FindAllSubCategories(id_categories, filter)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}

func (service *Controller) InsertCategoriesController(c echo.Context) error {
	categories := categories.CategoriesSpec{}
	c.Bind(&categories)

	id_admin_token := int(middleware.ExtractTokenKey(c, "id").(float64))

	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	err = service.categories_service.InsertCategories(
		categories,
		id_admin,
		id_admin_token)
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

	id_admin_token := int(middleware.ExtractTokenKey(c, "id").(float64))

	err = service.categories_service.ModifyCategories(
		categories,
		id_categories,
		id_admin,
		id_admin_token,
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
	id_admin_token := int(middleware.ExtractTokenKey(c, "id").(float64))

	err = service.categories_service.RemoveCategories(
		id_categories,
		id_admin,
		id_admin_token)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}
