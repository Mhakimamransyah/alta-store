package products

import (
	"altaStore/api/common"
	"altaStore/api/middleware"
	"altaStore/api/v1/products/request"

	"altaStore/business/products"
	productsimages "altaStore/business/products_images"
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	products_service       products.Service
	products_image_service productsimages.Service
}

func InitProductsController(service products.Service, products_images_s productsimages.Service) *Controller {
	return &Controller{
		products_service:       service,
		products_image_service: products_images_s,
	}
}

func (controller *Controller) CreateProductsController(c echo.Context) error {
	products_spec := products.ProductsSpec{}
	c.Bind(&products_spec)
	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	new_products, err := controller.products_service.InsertProducts(id_admin, products_spec, int(middleware.ExtractTokenKey(c, "id").(float64)))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	files := form.File["files"]
	file_name := strconv.Itoa(new_products.ID)

	products_images := productsimages.ProductImages{
		Products_ID: new_products.ID,
		Created_at:  time.Now(),
		FileName:    file_name,
	}

	err = controller.products_image_service.InsertNewImages(&products_images,
		files,
		int(middleware.ExtractTokenKey(c, "id").(float64)),
	)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	// }
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
	err = controller.products_service.ModifyProducts(
		id_admin,
		id_products,
		products_updatables,
		int(middleware.ExtractTokenKey(c, "id").(float64)),
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
	data, err := controller.products_service.DetailProducts(id_products)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}

func (controller *Controller) RemoveProductsController(c echo.Context) error {
	id_products, err := strconv.Atoi(c.Param("id_products"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	err = controller.products_service.RemoveProducts(id_products, id_admin)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) FindAllProductsController(c echo.Context) error {
	filter := request.NewFilterProducts(c)
	data, err := controller.products_service.FindAllProducts(*filter)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}

func (controller *Controller) RemoveProductsPictureController(c echo.Context) error {
	id_products, err := strconv.Atoi(c.Param("id_products"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	id_products_images, err := strconv.Atoi(c.Param("id_products_images"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	id_admin := int(middleware.ExtractTokenKey(c, "id").(float64))
	err = controller.products_image_service.RemoveProductsImages(id_admin, id_products, id_products_images, id_admin)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) InsertNewProductsPictureController(c echo.Context) error {
	id_products, err := strconv.Atoi(c.Param("id_products"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	id_admin := int(middleware.ExtractTokenKey(c, "id").(float64))

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	files := form.File["files"]

	products_images := productsimages.ProductImages{
		Products_ID: id_products,
		Created_at:  time.Now(),
		FileName:    strconv.Itoa(id_products),
	}

	err = controller.products_image_service.InsertNewImages(&products_images, files, id_admin)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
