package productsimages_test

import (
	"altaStore/business"
	productsimages "altaStore/business/products_images"
	"altaStore/business/products_images/mocks"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	created_by_id      = 1
	deleted_by_id      = 1
	id_admins          = 1
	id_products        = 1
	id_products_images = 1
	fileName           = "Filename.png"
	Path               = "path/Filename.png"
)

var (
	imagesFile              *multipart.FileHeader
	listImagesFile          []*multipart.FileHeader
	productsImageRepository mocks.Repository
	productsImagesService   productsimages.Service
	productsImagesData      productsimages.ProductImages
	listProductsImagesData  []productsimages.ProductImages
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestInsertNewImages(t *testing.T) {
	t.Run("Expects success insert new products images", func(t *testing.T) {
		productsImageRepository.On("CreateImages",
			mock.AnythingOfType("*productsimages.ProductImages"),
			imagesFile,
			created_by_id,
		).Return(nil).Once()
		err := productsImagesService.InsertNewImages(&productsImagesData, listImagesFile, created_by_id)
		assert.Nil(t, err)
	})

	t.Run("Expects failed insert new products images", func(t *testing.T) {
		productsImageRepository.On("CreateImages",
			mock.AnythingOfType("*productsimages.ProductImages"),
			imagesFile,
			created_by_id,
		).Return(business.ErrInternalServerError).Once()
		err := productsImagesService.InsertNewImages(&productsImagesData, listImagesFile, created_by_id)
		assert.Nil(t, err)
	})
}

func TestRemoveProductsImages(t *testing.T) {
	t.Run("Expects success remove products images", func(t *testing.T) {
		productsImageRepository.On("GetProductsImagesById", mock.AnythingOfType("int")).Return(&productsImagesData, nil).Once()
		productsImageRepository.On("DeleteProductsImages", &productsImagesData, deleted_by_id).Return(nil).Once()
		err := productsImagesService.RemoveProductsImages(id_admins, id_products, id_products_images, deleted_by_id)
		assert.Nil(t, err)
	})

	t.Run("Expects failed remove products images, not found", func(t *testing.T) {
		productsImageRepository.On("GetProductsImagesById", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()
		err := productsImagesService.RemoveProductsImages(id_admins, id_products, id_products_images, deleted_by_id)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed remove products images", func(t *testing.T) {
		productsImageRepository.On("GetProductsImagesById", mock.AnythingOfType("int")).Return(&productsImagesData, nil).Once()
		productsImageRepository.On("DeleteProductsImages", &productsImagesData, deleted_by_id).Return(business.ErrInternalServerError).Once()
		err := productsImagesService.RemoveProductsImages(id_admins, id_products, id_products_images, deleted_by_id)
		assert.NotNil(t, err)
	})
}

func setup() {
	productsImagesService = productsimages.InitProductsImagesService(&productsImageRepository)
	productsImagesData = productsimages.ProductImages{
		ID:          id_products_images,
		FileName:    fileName,
		Path:        Path,
		Products_ID: id_products,
	}
	listProductsImagesData = append(listProductsImagesData, productsImagesData)
	imagesFile = &multipart.FileHeader{}
	listImagesFile = append(listImagesFile, imagesFile)
}
