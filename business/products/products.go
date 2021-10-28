package products

import "time"

type Products struct {
	ID           int
	Stock        int
	Title        string
	Price        int
	Description  string
	Weight       float64
	Status       string
	AdminID      int
	CategoriesID int
	Created_at   time.Time
	Updated_at   time.Time
	Deleted_at   time.Time
}

func NewProducts(id_admin int, product_spec ProductsSpec) *Products {
	return &Products{
		Title:        product_spec.Title,
		Price:        product_spec.Price,
		Description:  product_spec.Description,
		Weight:       product_spec.Weight,
		Status:       "active",
		Stock:        product_spec.Stock,
		AdminID:      id_admin,
		CategoriesID: product_spec.Id_categories,
		Created_at:   time.Now(),
	}
}

func (old_products *Products) ModifyProducts(id_admin int, product_updatetables ProductsUpdatable) *Products {
	return &Products{
		ID:           old_products.ID,
		Stock:        product_updatetables.Stock,
		Title:        product_updatetables.Title,
		Price:        product_updatetables.Price,
		Description:  product_updatetables.Description,
		Weight:       product_updatetables.Weight,
		Status:       product_updatetables.Status,
		AdminID:      id_admin,
		CategoriesID: product_updatetables.CategoriesID,
		Updated_at:   time.Now(),
	}
}
