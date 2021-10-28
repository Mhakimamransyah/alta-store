package productsimages

import "time"

type ProductImages struct {
	ID          int
	Image_path  string
	Products_ID int
	Cretaed_at  time.Time
	Updated_at  time.Time
	Deleted_at  time.Time
}
