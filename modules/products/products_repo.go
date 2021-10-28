package products

import (
	"altaStore/business/products"
	"altaStore/modules/admins"
	"altaStore/modules/categories"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type ProductsTable struct {
	gorm.Model
	ID           int                        `gorm:"id;primaryKey:autoIncrement"`
	Stock        int                        `gorm:"stock;not null;default:0"`
	Title        string                     `gorm:"title;not null;uniqueIndex:category_name,sort:asc;type:varchar(100)"`
	Price        int                        `gorm:"price;not null;"`
	Description  string                     `gorm:"description;"`
	Weight       float64                    `gorm:"weight"`
	Status       string                     `gorm:"status;type:varchar(10)"`
	AdminID      int                        `gorm:"admin_id"`
	CategoriesID int                        `gorm:"categories_id"`
	Created_at   time.Time                  `gorm:"created_at;type:datetime;default:null"`
	Updated_at   time.Time                  `gorm:"updated_at;type:datetime;default:null"`
	Deleted_at   time.Time                  `gorm:"deleted_at;type:datetime;default:null"`
	Admin        admins.AdminsTable         `gorm:"foreignKey:AdminID"`
	Categories   categories.CategoriesTable `gorm:"foreignKey:CategoriesID"`
}

func ConvertProductsToProductsTable(products *products.Products) *ProductsTable {
	return &ProductsTable{
		ID:           products.ID,
		Stock:        products.Stock,
		Title:        products.Title,
		Price:        products.Price,
		Description:  products.Description,
		Weight:       products.Weight,
		Status:       products.Status,
		AdminID:      products.AdminID,
		CategoriesID: products.CategoriesID,
		Created_at:   products.Created_at,
		Updated_at:   products.Updated_at,
		Deleted_at:   products.Deleted_at,
	}
}

func ConvertProductsTableToProducts(products_table *ProductsTable) *products.Products {
	return &products.Products{
		ID:           products_table.ID,
		Stock:        products_table.Stock,
		Title:        products_table.Title,
		Price:        products_table.Price,
		Description:  products_table.Description,
		Weight:       products_table.Weight,
		Status:       products_table.Status,
		AdminID:      products_table.AdminID,
		CategoriesID: products_table.CategoriesID,
		Created_at:   products_table.Created_at,
		Updated_at:   products_table.Updated_at,
		Deleted_at:   products_table.Deleted_at,
	}
}

func InitProducstRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		DB: db,
	}
}

func (repository *GormRepository) GetAllProducts(limit, offset int) (*[]products.Products, error) {
	var list_products_table []ProductsTable
	err := repository.DB.Where("status = ?", "active").Offset(offset - 1).Limit(limit).Find(&list_products_table).Error
	if err != nil {
		return nil, err
	}

	var list_products []products.Products
	for _, data := range list_products_table {
		list_products = append(list_products, *ConvertProductsTableToProducts(&data))
	}
	return &list_products, nil
}

func (repository *GormRepository) GetDetailProducts(id_products int) (*products.Products, error) {
	products_table := ProductsTable{}
	err := repository.DB.Where("id = ?", id_products).First(&products_table).Error
	if err != nil {
		return nil, err
	}
	products := ConvertProductsTableToProducts(&products_table)
	return products, nil
}

func (repository *GormRepository) CreateProducts(products *products.Products, createdBy string) error {
	products_table := ConvertProductsToProductsTable(products)
	err := repository.DB.Save(products_table).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *GormRepository) UpdateProducts(id_products int, products *products.Products, modifiedBy string) error {
	products_table := ConvertProductsToProductsTable(products)
	err := repository.DB.Where("id = ?", id_products).Model(products_table).Updates(ProductsTable{
		Stock:        products.Stock,
		Title:        products.Title,
		Price:        products.Price,
		Description:  products.Description,
		Weight:       products.Weight,
		Status:       products.Status,
		CategoriesID: products.CategoriesID,
		Updated_at:   products.Updated_at,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
