package categories

import (
	"altaStore/business/categories"
	"altaStore/modules/admins"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type CategoriesTable struct {
	gorm.Model
	ID         int                `gorm:"id;primaryKey:autoIncrement"`
	Name       string             `gorm:"name;type:varchar(20);uniqueIndex:category_name,sort:asc;not null"`
	AdminID    int                `gorm:"admin_id;not null"`
	Status     string             `gorm:"status;not null"`
	ParentID   int                `gorm:"parent_id;default:null"`
	Created_at time.Time          `gorm:"created_at;type:datetime;default:null"`
	Updated_at time.Time          `gorm:"updated_at;type:datetime;default:null"`
	Deleted_at time.Time          `gorm:"deleted_at;type:datetime;default:null"`
	Admin      admins.AdminsTable `gorm:"foreignKey:AdminID"`
}

func ConvertCategoriesToCategoriesTable(categories *categories.Categories) *CategoriesTable {
	return &CategoriesTable{
		ID:         categories.ID,
		Name:       categories.Name,
		Status:     categories.Status,
		AdminID:    categories.AdminID,
		ParentID:   categories.Parent_id,
		Created_at: categories.Created_at,
		Updated_at: categories.Updated_at,
		Deleted_at: categories.Deleted_at,
	}
}

func ConvertCategoriesTableToCategories(categories_table *CategoriesTable) *categories.Categories {
	return &categories.Categories{
		ID:         categories_table.ID,
		Name:       categories_table.Name,
		Status:     categories_table.Status,
		Parent_id:  categories_table.ParentID,
		AdminID:    categories_table.AdminID,
		Created_at: categories_table.Created_at,
		Updated_at: categories_table.Updated_at,
		Deleted_at: categories_table.Deleted_at,
	}
}

func InitCategoriesRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		DB: db,
	}
}

func (repository *GormRepository) CreateCategories(categories *categories.Categories, createdBy int) error {
	categories_table := ConvertCategoriesToCategoriesTable(categories)
	err := repository.DB.Save(categories_table).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *GormRepository) UpdateCategories(categories *categories.Categories, modifiedBy int) error {
	categories_table := ConvertCategoriesToCategoriesTable(categories)
	err := repository.DB.Where("ID = ?", categories.ID).Model(&categories_table).Updates(CategoriesTable{
		Name:       categories.Name,
		Status:     categories.Status,
		Updated_at: categories.Updated_at,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *GormRepository) GetCategories(categories_search *categories.FilterCategories) (*[]categories.Categories, error) {
	var list_categories_tables []CategoriesTable
	model := repository.setCategoriesFilter(categories_search).Where("parent_id IS NULL")
	err := model.Find(&list_categories_tables).Error
	if err != nil {
		return nil, err
	}
	var list_categories []categories.Categories
	for _, data := range list_categories_tables {
		list_categories = append(list_categories, *repository.CountChildCategories(
			*ConvertCategoriesTableToCategories(&data),
			data,
		))
	}
	return &list_categories, nil
}

func (repository *GormRepository) CountChildCategories(categories categories.Categories, categories_table CategoriesTable) *categories.Categories {
	var count_child int64
	repository.DB.Where("parent_id = ? AND status = ?", categories.ID, "active").Model(&categories_table).Count(&count_child)
	categories.Count_child_categories = int(count_child)
	return &categories
}

func (repository *GormRepository) GetSubCategories(id_categories int, categories_search *categories.FilterCategories) (*[]categories.Categories, error) {
	var list_categories_tables []CategoriesTable
	model := repository.setCategoriesFilter(categories_search).Where("parent_id = ?", id_categories)
	err := model.Find(&list_categories_tables).Error
	if err != nil {
		return nil, err

	}
	var list_categories []categories.Categories
	for _, data := range list_categories_tables {
		list_categories = append(list_categories,
			*repository.CountChildCategories(*ConvertCategoriesTableToCategories(&data),
				data))
	}
	return &list_categories, nil
}

func (repository *GormRepository) DeleteCategories(id_categories int, deletedBy int) error {
	err := repository.DB.Where("ID = ?", id_categories).Delete(&CategoriesTable{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *GormRepository) GetCategoriesById(id_categories int) (*categories.Categories, error) {
	var categories_table CategoriesTable
	err := repository.DB.Where("ID = ?", id_categories).First(&categories_table).Error
	if err != nil {
		return nil, err
	}
	return ConvertCategoriesTableToCategories(&categories_table), nil
}

func (repository *GormRepository) setCategoriesFilter(filter_Categories *categories.FilterCategories) *gorm.DB {
	model := repository.DB
	if filter_Categories.AdminId != 0 {
		model = model.Where("admin_id = ?", filter_Categories.AdminId)
	}
	if filter_Categories.Query != "" {
		model = model.Where("name LIKE ?", filter_Categories.Query+"%")
	}

	if filter_Categories.Status != "active" {
		model = model.Where("status != ?", "active")
	} else {
		model = model.Where("status = ?", "active")
	}

	if filter_Categories.SortName == "asc" || filter_Categories.SortName == "desc" {
		model = model.Order("name " + filter_Categories.SortName)
	}

	if filter_Categories.SortDate != "asc" || filter_Categories.SortDate == "desc" {
		model = model.Order("created_at " + filter_Categories.SortDate)
	}

	if filter_Categories.Offset != 0 {
		model = model.Offset(filter_Categories.Offset - 1)
	} else {
		model = model.Offset(0)
	}

	if filter_Categories.Limit != 0 {
		model = model.Limit(filter_Categories.Limit)
	} else {
		model = model.Limit(100)
	}

	return model
}
