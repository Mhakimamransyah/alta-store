package api

import (
	"altaStore/api/middleware"
	"altaStore/api/v1/address"
	"altaStore/api/v1/admins"
	"altaStore/api/v1/auth"
	"altaStore/api/v1/cart"
	"altaStore/api/v1/categories"
	"altaStore/api/v1/products"
	"altaStore/api/v1/transaction"
	"altaStore/api/v1/user"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

//RegisterPath Register all API with routing path
func RegisterPath(
	e *echo.Echo,
	authController *auth.Controller,
	userController *user.Controller,
	addressController *address.Controller,
	cartController *cart.Controller,
	adminController *admins.Controller,
	categoriesController *categories.Controller,
	productsController *products.Controller,
	transactionController *transaction.Controller) {
	// if authController == nil || userController == nil || addressController == nil || cartController == nil {
	// 	panic("Controller parameter cannot be nil")
	// }

	// 	//authentication with Versioning endpoint
	authV1 := e.Group("v1/auth")
	authV1.POST("/login", authController.Login)
	authV1.POST("/register", userController.InsertUser)

	//user with Versioning endpoint
	userV1 := e.Group("v1/users")
	userV1.Use(middleware.JWTMiddleware())
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUser)
	userV1.POST("/address", addressController.InsertAddress)
	userV1.GET("/address", addressController.GetAllAddress)
	userV1.GET("/address/default", addressController.GetDefaultAddress)

	cartV1 := e.Group("v1/cart")
	cartV1.Use(middleware.JWTMiddleware())
	cartV1.POST("", cartController.AddToCart)
	cartV1.GET("", cartController.GetActiveCart)

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

	// Transaction
	transaction := e.Group("/v1/transaction")
	transaction.Use(middleware.JWTMiddleware())
	transaction.POST("", transactionController.Checkout)
	transaction.GET("", transactionController.GetAllTransaction)
	transaction.GET("/invoice", transactionController.GetTransactionByInvoice)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
