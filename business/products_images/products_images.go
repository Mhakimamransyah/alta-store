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
