package api

import (
	"altaStore/api/middleware"
	"altaStore/api/v1/admins"
	"altaStore/api/v1/categories"
	"altaStore/api/v1/products"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	adminController *admins.Controller,
	categoriesController *categories.Controller,
	productsController *products.Controller,
) {

	e.POST("admins/login", adminController.LoginController)
	e.POST("adminmockdata", adminController.CreateAdminController)

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

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
