package main

import (
	"altaStore/api"
	"altaStore/config"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	addressController "altaStore/api/v1/address"
	AdminsController "altaStore/api/v1/admins"
	authController "altaStore/api/v1/auth"
	cartController "altaStore/api/v1/cart"
	CategoriesController "altaStore/api/v1/categories"
	ProductsController "altaStore/api/v1/products"
	transactionController "altaStore/api/v1/transaction"
	userController "altaStore/api/v1/user"
	addressService "altaStore/business/address"
	AdminsService "altaStore/business/admins"
	authService "altaStore/business/auth"
	cartService "altaStore/business/cart"
	CategoriesService "altaStore/business/categories"
	ProductsService "altaStore/business/products"
	productsimages "altaStore/business/products_images"
	transactionService "altaStore/business/transaction"
	userService "altaStore/business/user"
	addressRepository "altaStore/modules/address"
	AdminsRepository "altaStore/modules/admins"
	cartRepository "altaStore/modules/cart"
	CategoriesRepository "altaStore/modules/categories"
	migration "altaStore/modules/migration"
	ProductsRepository "altaStore/modules/products"
	ProductsImages "altaStore/modules/products_images"
	transactionRepository "altaStore/modules/transaction"
	userRepository "altaStore/modules/user"
	utilService "altaStore/util/password"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": os.Getenv("ALTASTORE_DB_USERNAME"),
		"DB_Password": os.Getenv("ALTASTORE_DB_PASSWORD"),
		"DB_Port":     os.Getenv("ALTASTORE_DB_PORT"),
		"DB_Host":     os.Getenv("ALTASTORE_DB_ADDRESS"),
		"DB_Name":     os.Getenv("ALTASTORE_DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)
	fmt.Println("Connect to ", db.Migrator().CurrentDatabase())
	return db
}

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	dbConnection := newDatabaseConnection(config)

	utilService := utilService.NewUtil()
	userRepo := userRepository.NewGormDBRepository(dbConnection)
	userService := userService.NewService(userRepo, utilService)
	userController := userController.NewController(userService)

	authService := authService.NewService(userService, utilService)
	authController := authController.NewController(authService)

	addressRepo := addressRepository.NewGormDBRepository(dbConnection)
	addressService := addressService.NewService(addressRepo)
	addressController := addressController.NewController(addressService)

	adminRepository := AdminsRepository.InitAdminRepository(dbConnection)
	adminService := AdminsService.InitAdminService(adminRepository)
	adminController := AdminsController.InitAdminController(adminService)

	categoriesRepository := CategoriesRepository.InitCategoriesRepository(dbConnection)
	categoriesService := CategoriesService.InitCategoriesService(categoriesRepository)
	categoriesController := CategoriesController.InitCategoriesController(categoriesService)

	productsImagesRepository := ProductsImages.InitProductsImagesRepository(dbConnection)
	productsImagesService := productsimages.InitProductsImagesService(productsImagesRepository)

	productsRepository := ProductsRepository.InitProducstRepository(dbConnection)
	productsService := ProductsService.InitProductsService(productsRepository, productsImagesRepository)
	productsController := ProductsController.InitProductsController(productsService, productsImagesService)

	cartRepo := cartRepository.NewGormDBRepository(dbConnection)
	cartService := cartService.NewService(cartRepo, productsRepository)
	cartController := cartController.NewController(cartService)

	transactionRepo := transactionRepository.NewGormDBRepository(dbConnection)
	transactionService := transactionService.NewService(transactionRepo, cartRepo, addressRepo, productsRepository)
	transactionController := transactionController.NewController(transactionService)

	//create echo http
	e := echo.New()

	//register API path and handler
	api.RegisterPath(
		e,
		authController,
		userController,
		addressController,
		cartController,
		adminController,
		categoriesController,
		productsController,
		transactionController,
	)

	// run server
	go func() {
		address := fmt.Sprintf(":%d", config.AppPort)
		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
