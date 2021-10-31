package products

type Service interface {
	FindAllProducts(filter FilterProducts) (*[]Products, error)
	DetailProducts(id_products int) (*Products, error)
	InsertProducts(id_admin int, products_spec ProductsSpec, createdById int) (*Products, error)
	ModifyProducts(id_admin, id_products int, products_updatable ProductsUpdatable, modifiedById int) error
	RemoveProducts(id_products int, deletedById int) error
}

type Repository interface {
	GetAllProducts(filter FilterProducts) (*[]Products, error)
	GetDetailProducts(id_products int) (*Products, error)
	CreateProducts(products *Products, createdById int) (*Products, error)
	UpdateProducts(id_products int, products *Products, modifiedById int) error
	DeleteProducts(products *Products) error
	UpdateStocks(id_products, value int, operation string) error
}
