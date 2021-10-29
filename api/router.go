package api

import (
	"altaStore/api/middleware"
	"altaStore/api/v1/admins"
	"altaStore/api/v1/categories"
	"altaStore/api/v1/products"

	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	adminController *admins.Controller,
	categoriesController *categories.Controller,
	productsController *products.Controller,
) {

	e.POST("admins/login", adminController.LoginController)
	e.POST("adminmockdata", adminController.CreateMockAdminController)

	// Admin
	admin := e.Group("admins")
	admin.Use(middleware.JWTMiddleware())
	admin.GET("", adminController.GetAdminController)
	admin.GET("/:username", adminController.GetAdminByUsername)
	admin.POST("", adminController.CreateAdminController)
	admin.PUT("/:username", adminController.ModifyAdminController)

	admin.POST("/:id_admin/categories", categoriesController.InsertCategoriesController)
	admin.PUT("/:id_admin/categories/:id_categories", categoriesController.ModifyCategoriesController)
	admin.DELETE("/:id_admin/categories/:id_categories", categoriesController.RemoveCategoriesController)

	admin.POST("/:id_admin/products", productsController.CreateProductsController)
	admin.PUT("/:id_admin/products/:id_products", productsController.ModifyProductsController)

	// Categories
	categories := e.Group("categories")
	categories.Use(middleware.JWTMiddleware())
	categories.GET("", categoriesController.GetAllCategoriesController)
	categories.GET("/:id_categories", categoriesController.GetAllSubCategoriesController)

	// Products
	products := e.Group("products")
	products.Use(middleware.JWTMiddleware())
	products.GET("", productsController.FindAllProductsController)
	products.GET("/:id_products", productsController.DetailProductsController)
	products.POST("/:id_products/products_images", productsController.InsertNewProductsPictureController)
	products.DELETE("/:id_products/products_images/:id_products_images", productsController.RemoveProductsPictureController)
	products.GET("/search", productsController.FindAllProductsController)

	// Products Images
	fs := http.FileServer(http.Dir("products_image/"))
	e.GET("/products_img/*", echo.WrapHandler(http.StripPrefix("/products_img/", fs)))

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
