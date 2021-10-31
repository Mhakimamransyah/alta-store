package products_test

import (
	"altaStore/business"
	"altaStore/business/products"
	"altaStore/business/products/mocks"
	products_images "altaStore/business/products_images"
	productsimages "altaStore/business/products_images/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

const (
	ProductsID    = 1
	Title         = "products title"
	Price         = 10000
	Description   = "products descriptions"
	Weight        = 3.3
	CategoriesID  = 1
	Stock         = 20
	Status        = "active"
	AdminID       = 5
	created_by_id = 10
)

var (
	producstSpec             products.ProductsSpec
	productsService          products.Service
	productsRepository       mocks.Repository
	productsImagesRepository productsimages.Repository
	productsUpdate           products.ProductsUpdatable
	productsData             products.Products
	InvalidProductsData      products.ProductsSpec
	productsFilter           products.FilterProducts
	listProducts             []products.Products
	listProductsImages       []products_images.ProductImages
	productsImagesData       products_images.ProductImages
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAllProducts(t *testing.T) {
	t.Run("Expects success find all products", func(t *testing.T) {
		productsRepository.On("GetAllProducts", productsFilter).Return(&listProducts, nil).Once()
		productsImagesRepository.On("GetListProductsImagesByIdProducts", ProductsID).Return(&listProductsImages, nil)
		list_products, err := productsService.FindAllProducts(productsFilter)
		assert.NotNil(t, list_products)
		assert.Nil(t, err)
	})

	t.Run("Expects failed find all products", func(t *testing.T) {
		productsRepository.On("GetAllProducts", productsFilter).Return(nil, business.ErrNotFound).Once()
		list_products, err := productsService.FindAllProducts(productsFilter)
		assert.Nil(t, list_products)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed find all products images", func(t *testing.T) {
		productsRepository.On("GetAllProducts", productsFilter).Return(&listProducts, nil).Once()
		productsImagesRepository.On("GetListProductsImagesByIdProducts", ProductsID).Return(nil, business.ErrNotFound)
		list_products, err := productsService.FindAllProducts(productsFilter)
		assert.NotNil(t, list_products)
		assert.Nil(t, err)
	})
}

func TestDetailProducts(t *testing.T) {
	t.Run("Expects success find detail products", func(t *testing.T) {
		productsRepository.On("GetDetailProducts", ProductsID).Return(&productsData, nil).Once()
		productsImagesRepository.On("GetListProductsImagesByIdProducts", ProductsID).Return(&listProductsImages, nil)
		products, err := productsService.DetailProducts(ProductsID)
		assert.NotNil(t, products)
		assert.Nil(t, err)
	})

	t.Run("Expects failed find detail products", func(t *testing.T) {
		productsRepository.On("GetDetailProducts", ProductsID).Return(nil, business.ErrNotFound).Once()
		products, err := productsService.DetailProducts(ProductsID)
		assert.Nil(t, products)
		assert.NotNil(t, err)
	})
}

func TestInsertProducts(t *testing.T) {
	t.Run("Expects success insert products", func(t *testing.T) {
		productsRepository.On("CreateProducts", mock.AnythingOfType("*products.Products"), AdminID).Return(&productsData, nil).Once()
		products, err := productsService.InsertProducts(AdminID, producstSpec, AdminID)
		assert.NotNil(t, products)
		assert.Nil(t, err)
	})

	t.Run("Expects failed insert products", func(t *testing.T) {
		productsRepository.On("CreateProducts", mock.AnythingOfType("*products.Products"), AdminID).Return(nil, business.ErrInternalServerError).Once()
		products, err := productsService.InsertProducts(AdminID, producstSpec, AdminID)
		assert.Nil(t, products)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed insert products, invalid spec", func(t *testing.T) {
		productsRepository.On("CreateProducts", mock.AnythingOfType("*products.Products"), AdminID).Return(nil, business.ErrInternalServerError).Once()
		products, err := productsService.InsertProducts(AdminID, InvalidProductsData, AdminID)
		assert.Nil(t, products)
		assert.NotNil(t, err)
	})
}

func TestModifyProducts(t *testing.T) {
	t.Run("Expects success update products", func(t *testing.T) {
		productsRepository.On("GetDetailProducts", ProductsID).Return(&productsData, nil).Once()
		productsRepository.On("UpdateProducts", ProductsID, mock.AnythingOfType("*products.Products"), AdminID).Return(nil).Once()
		err := productsService.ModifyProducts(AdminID, ProductsID, productsUpdate, AdminID)
		assert.Nil(t, err)
	})

	t.Run("Expects failed update products", func(t *testing.T) {
		productsRepository.On("GetDetailProducts", ProductsID).Return(&productsData, nil).Once()
		productsRepository.On("UpdateProducts", ProductsID, mock.AnythingOfType("*products.Products"), AdminID).Return(business.ErrInternalServerError).Once()
		err := productsService.ModifyProducts(AdminID, ProductsID, productsUpdate, AdminID)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed update products, products not found", func(t *testing.T) {
		productsRepository.On("GetDetailProducts", ProductsID).Return(nil, business.ErrNotFound).Once()
		err := productsService.ModifyProducts(AdminID, ProductsID, productsUpdate, AdminID)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed update products, unauthorized admin", func(t *testing.T) {
		productsRepository.On("GetDetailProducts", ProductsID).Return(&productsData, nil).Once()
		err := productsService.ModifyProducts(100, ProductsID, productsUpdate, AdminID)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed update products, data products invalid", func(t *testing.T) {
		productsRepository.On("GetDetailProducts", ProductsID).Return(&productsData, nil).Once()
		err := productsService.ModifyProducts(100, ProductsID, products.ProductsUpdatable{}, AdminID)
		assert.NotNil(t, err)
	})
}

func TestRemoveProducts(t *testing.T) {

}

func setup() {
	productsData = products.Products{
		ID:           ProductsID,
		Stock:        Stock,
		Title:        Title,
		Price:        Price,
		Description:  Description,
		Weight:       Weight,
		Status:       Status,
		AdminID:      AdminID,
		CategoriesID: CategoriesID,
	}

	InvalidProductsData = products.ProductsSpec{
		Stock:       Stock,
		Description: Description,
		Weight:      Weight,
	}

	producstSpec = products.ProductsSpec{
		Title:         Title,
		Price:         Price,
		Description:   Description,
		Weight:        Weight,
		Id_categories: CategoriesID,
		Stock:         Stock,
	}

	productsUpdate = products.ProductsUpdatable{
		Title:        Title,
		Price:        Price,
		Description:  Description,
		Weight:       Weight,
		CategoriesID: CategoriesID,
		Stock:        Stock,
		Status:       Status,
	}

	productsFilter = products.FilterProducts{
		CategoriesId: CategoriesID,
		Sort:         "asc",
		Page:         1,
		Per_page:     100,
	}

	productsImagesData = products_images.ProductImages{
		ID:          1,
		FileName:    "filename.png",
		Path:        "path/filename.png",
		Products_ID: ProductsID,
	}

	listProductsImages = append(listProductsImages, productsImagesData)

	listProducts = append(listProducts, productsData)

	productsService = products.InitProductsService(&productsRepository, &productsImagesRepository)
}
