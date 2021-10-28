package products

type Service interface {
	FindAllProducts(limit, offset int) (*[]Products, error)
	DetailProducts(id_products int) (*Products, error)
	InsertProducts(id_admin int, products_spec ProductsSpec, createdBy string) error
	ModifyProducts(id_admin, id_products int, products_updatable ProductsUpdatable, modifiedBy string) error
}

type Repository interface {
	GetAllProducts(limit, offset int) (*[]Products, error)
	GetDetailProducts(id_products int) (*Products, error)
	CreateProducts(products *Products, createdBy string) error
	UpdateProducts(id_products int, products *Products, modifiedBy string) error
}
