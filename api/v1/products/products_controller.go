package products

import (
	"altaStore/api/common"
	"altaStore/api/middleware"
	"altaStore/business/products"
	productsimages "altaStore/business/products_images"
	"fmt"
	"io"
	"os"
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
	mockService(c)
	c.Bind(&products_spec)
	id_admin, err := strconv.Atoi(c.Param("id_admin"))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	new_products, err := controller.products_service.InsertProducts(id_admin, products_spec, middleware.ExtractTokenKey(c, "username").(string))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	files := form.File["files"]
	file_name := strconv.Itoa(new_products.ID)

	products_images := productsimages.ProductImages{
		Products_ID: new_products.ID,
		Created_at:  time.Now(),
		FileName:    file_name,
	}

	err = controller.products_image_service.InsertNewImages(products_images,
		files,
		middleware.ExtractTokenKey(c, "username").(string))
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}

	// }
	return c.JSON(common.NewSuccessResponseWithoutData())
}

func mockService(c echo.Context) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	// data, err := c.FormFile("foto")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("Ukuran file : ", float64(data.Size/1024/1024))
	// fmt.Println("Type file : ", data.Header["Content-Type"][0])
	// src, err := data.Open()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer src.Close()
	// // Destination
	// dst, err := os.Create(data.Filename)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer dst.Close()
	// if _, err = io.Copy(dst, src); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("done")

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			fmt.Println(err)
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			fmt.Println(err)
		}

	}

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
	data, err := controller.products_service.DetailProducts(id_products)
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
	data, err := controller.products_service.FindAllProducts(limit, offset)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	return c.JSON(common.NewSuccessResponse(data))
}
