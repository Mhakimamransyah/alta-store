package migration

import (
	"altaStore/modules/admins"
	"altaStore/modules/categories"
	"altaStore/modules/products"
	productsimages "altaStore/modules/products_images"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&admins.AdminsTable{}, &categories.CategoriesTable{}, &products.ProductsTable{}, &productsimages.ProductsImagesTable{})
}
