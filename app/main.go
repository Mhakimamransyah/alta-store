package main

import (
	"altaStore/api"
	"altaStore/config"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	userController "altaStore/api/v1/user"
	userService "altaStore/business/user"
	migration "altaStore/modules/migration"
	userRepository "altaStore/modules/user"

	"github.com/labstack/echo"
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

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	config.DbUsername,
	// 	config.DbPassword,
	// 	config.DbAddress,
	// 	strconv.Itoa(config.DbPort),
	// 	config.DbName)

	// fmt.Println(connectionString)

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
	fmt.Println(config)

	//initialize database connection based on given config
	dbConnection := newDatabaseConnection(config)

	//initiate user repository
	userRepo := userRepository.NewGormDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(userRepo)

	//initiate user controller
	userController := userController.NewController(userService)

	// //initiate pet repository
	// petRepo := petRepository.NewGormDBRepository(dbConnection)

	// //initiate pet service
	// petService := petService.NewService(petRepo)

	// //initiate pet controller
	// petController := petController.NewController(petService)

	//initiate auth service
	// authService := authService.NewService(userService)

	//initiate auth controller
	// authController := authController.NewController(authService)

	//create echo http
	e := echo.New()

	//register API path and handler
	api.RegisterPath(e, userController)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%d", config.AppPort)

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
