package productsimages

import (
	"altaStore/modules/products"
	"time"
)

type ProductsImagesTable struct {
	ID          int                    `gorm:"id;primaryKey:autoIncrement"`
	Image_path  string                 `gorm:"image_path"`
	Products_ID int                    `gorm:"id_products"`
	Cretaed_at  time.Time              `gorm:"created_at;type:datetime;default:null"`
	Updated_at  time.Time              `gorm:"updated_at;type:datetime;default:null"`
	Deleted_at  time.Time              `gorm:"deleted_at;type:datetime;default:null"`
	Products    products.ProductsTable `gorm:"foreignKey:Products_ID"`
}
