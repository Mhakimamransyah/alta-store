package products

import (
	"altaStore/api/common"
	"altaStore/api/middleware"
	"altaStore/business/products"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service products.Service
}

func InitProductsController(service products.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) CreateProductsController(c echo.Context) error {
	products_spec := products.ProductsSpec{}
	c.Bind(&products_spec)
	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	err = controller.service.InsertProducts(id_admin, products_spec, middleware.ExtractTokenKey(c, "username").(string))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) ModifyProductsController(c echo.Context) error {
	products_updatables := products.ProductsUpdatable{}
	c.Bind(&products_updatables)
	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	id_products, err := strconv.Atoi(c.Param("id_products"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	err = controller.service.ModifyProducts(
		id_admin,
		id_products,
		products_updatables,
		middleware.ExtractTokenKey(c, "username").(string),
	)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) DetailProductsController(c echo.Context) error {
	id_products, err := strconv.Atoi(c.Param("id_products"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	data, err := controller.service.DetailProducts(id_products)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}

func (controller *Controller) FindAllProductsController(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Invalid per_page type"))
	}
	offset, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage("Invalid page type"))
	}
	data, err := controller.service.FindAllProducts(limit, offset)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}
