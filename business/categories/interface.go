package categories

type Service interface {
	FindAllCategories(limit, offset int) (*[]Categories, error)
	FindAllSubCategories(id_categories, limit, offset int) (*[]Categories, error)
	InsertCategories(categories_spec CategoriesSpec, id_admin int, createdBy string) error
	ModifyCategories(categories_updatable CategoriesUpdatable, id_categories int, id_admin int, modifiedBy string) error
	RemoveCategories(id_categories int, id_admin int, deletedBy string) error
}

type Repository interface {
	GetCategories(limit, offset int) (*[]Categories, error)
	GetSubCategories(id_categories, limit, offset int) (*[]Categories, error)
	GetCategoriesById(id_categories int) (*Categories, error)
	CreateCategories(categories *Categories, createdBy string) error
	UpdateCategories(categories *Categories, modifiedBy string) error
	DeleteCategories(id_categories int, deletedBy string) error
}
