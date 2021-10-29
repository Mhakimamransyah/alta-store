package productsimages

import (
	"time"
)

type ProductImages struct {
	ID          int
	FileName    string
	Path        string
	Products_ID int
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_at  time.Time
}

func NewProductsImage(id_products int, image_path string) *ProductImages {
	return &ProductImages{
		Products_ID: id_products,
		Created_at:  time.Now(),
	}
}
