package migration

import (
	"altaStore/modules/address"
	"altaStore/modules/admins"
	"altaStore/modules/cart"
	"altaStore/modules/categories"
	"altaStore/modules/products"
	productsimages "altaStore/modules/products_images"
	"altaStore/modules/transaction"
	"altaStore/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&address.Address{},
		&cart.Cart{},
		&cart.CartDetail{},
		&admins.AdminsTable{},
		&categories.CategoriesTable{},
		&products.ProductsTable{},
		&productsimages.ProductsImagesTable{},
		&transaction.Transaction{},
	)
}
