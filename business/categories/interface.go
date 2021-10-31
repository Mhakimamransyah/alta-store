package categories

type Service interface {
	FindAllCategories(categories_search *FilterCategories) (*[]Categories, error)
	FindAllSubCategories(id_categories int, categories_search *FilterCategories) (*[]Categories, error)
	InsertCategories(categories_spec CategoriesSpec, id_admin int, createdById int) error
	ModifyCategories(categories_updatable CategoriesUpdatable, id_categories int, id_admin int, modifiedById int) error
	RemoveCategories(id_categories int, id_admin int, deletedById int) error
}

type Repository interface {
	GetCategories(categories_search *FilterCategories) (*[]Categories, error)
	GetSubCategories(id_categories int, categories_search *FilterCategories) (*[]Categories, error)
	GetCategoriesById(id_categories int) (*Categories, error)
	CreateCategories(categories *Categories, createdById int) error
	UpdateCategories(categories *Categories, modifiedById int) error
	DeleteCategories(id_categories int, deletedById int) error
}
