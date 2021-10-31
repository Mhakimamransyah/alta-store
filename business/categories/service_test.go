package categories_test

import (
	"altaStore/business"
	"altaStore/business/categories"
	"altaStore/business/categories/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// categoriesMock "altaStore/business/categories/mocks"
)

const (
	Id             = 1
	CategoriesName = "categories_name"
	ParentId       = 2
	AdminId        = 1
	Password       = "12345"
	Created_by     = "created_by"
	Created_by_id  = 1
	Status         = "active"
	DeletedBy      = 2
	limit          = 10
	offset         = 1
)

var (
	categoiresService    categories.Service
	categoriesSpec       categories.CategoriesSpec
	categoriesUpdate     categories.CategoriesUpdatable
	categoriesRepository mocks.Repository
	listCategories       []categories.Categories
	categoriesData       categories.Categories
	filterCategories     categories.FilterCategories
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAllCategories(t *testing.T) {
	t.Run("Expects success get all categories", func(t *testing.T) {
		categoriesRepository.On("GetCategories", &filterCategories).Return(&listCategories, nil).Once()
		list_categories, err := categoiresService.FindAllCategories(&filterCategories)
		assert.NotNil(t, list_categories)
		assert.Nil(t, err)
	})

	t.Run("Expect failed get all categories", func(t *testing.T) {
		categoriesRepository.On("GetCategories", &filterCategories).Return(nil, business.ErrNotFound).Once()
		list_categories, err := categoiresService.FindAllCategories(&filterCategories)
		assert.NotNil(t, err)
		assert.Nil(t, list_categories)
	})
}

func TestFindAllSubCategories(t *testing.T) {
	t.Run("Expects success get all sub categories", func(t *testing.T) {
		categoriesRepository.On("GetSubCategories", Id, &filterCategories).Return(&listCategories, nil).Once()
		list_categories, err := categoiresService.FindAllSubCategories(Id, &filterCategories)
		assert.NotNil(t, list_categories)
		assert.Nil(t, err)
	})

	t.Run("Expects failed get all sub categories", func(t *testing.T) {
		categoriesRepository.On("GetSubCategories", Id, &filterCategories).Return(nil, business.ErrNotFound).Once()
		list_sub_categories, err := categoiresService.FindAllSubCategories(Id, &filterCategories)
		assert.NotNil(t, err)
		assert.Nil(t, list_sub_categories)
	})
}

func TestInsertCategories(t *testing.T) {
	t.Run("Expects success insert sub categories", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", categoriesSpec.ParentID).Return(&categories.Categories{
			ID:        categoriesSpec.ParentID,
			Name:      "Parent categories",
			Parent_id: 0,
		}, nil).Once()
		categoriesRepository.On("CreateCategories", mock.AnythingOfType("*categories.Categories"), AdminId).Return(nil).Once()
		err := categoiresService.InsertCategories(categoriesSpec, AdminId, AdminId)
		assert.Nil(t, err)
	})

	t.Run("Expects failed insert categories", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", categoriesSpec.ParentID).Return(&categories.Categories{
			ID:        categoriesSpec.ParentID,
			Name:      "Parent categories",
			Parent_id: 0,
		}, nil).Once()
		categoriesRepository.On("CreateCategories", mock.AnythingOfType("*categories.Categories"), AdminId).Return(business.ErrInternalServerError).Once()
		err := categoiresService.InsertCategories(categoriesSpec, AdminId, AdminId)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed insert categories, parent not found", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", categoriesSpec.ParentID).Return(nil, business.ErrNotFound).Once()
		err := categoiresService.InsertCategories(categoriesSpec, AdminId, AdminId)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed insert categories, Invalid spec", func(t *testing.T) {
		err := categoiresService.InsertCategories(categories.CategoriesSpec{
			ParentID: ParentId,
			AdminID:  AdminId,
		}, AdminId, AdminId)
		assert.NotNil(t, err)
	})
}

func TestModifyCategories(t *testing.T) {
	t.Run("Expects success modify categories", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(&categoriesData, nil).Once()
		categoriesRepository.On("UpdateCategories", mock.AnythingOfType("*categories.Categories"), AdminId).Return(nil).Once()
		err := categoiresService.ModifyCategories(categoriesUpdate, Id, AdminId, AdminId)
		assert.Nil(t, err)
	})

	t.Run("Expects failed modify categories, categories not found", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(nil, business.ErrNotFound).Once()
		err := categoiresService.ModifyCategories(categoriesUpdate, Id, AdminId, AdminId)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed modify categories, invalid spec", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(&categoriesData, nil).Once()
		err := categoiresService.ModifyCategories(categories.CategoriesUpdatable{
			Status: "non active",
		}, Id, AdminId, AdminId)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed modify categories, unauthorized", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(&categories.Categories{
			AdminID: 100,
		}, nil).Once()
		err := categoiresService.ModifyCategories(categoriesUpdate, Id, AdminId, AdminId)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed modify categories", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(&categoriesData, nil).Once()
		categoriesRepository.On("UpdateCategories", mock.AnythingOfType("*categories.Categories"), AdminId).Return(business.ErrHasBeenModified).Once()
		err := categoiresService.ModifyCategories(categoriesUpdate, Id, AdminId, AdminId)
		assert.NotNil(t, err)
	})
}

func TestRemoveCategories(t *testing.T) {
	t.Run("Expects success delete categories", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(&categoriesData, nil).Once()
		categoriesRepository.On("DeleteCategories", Id, DeletedBy).Return(nil).Once()
		err := categoiresService.RemoveCategories(Id, AdminId, DeletedBy)
		assert.Nil(t, err)
	})

	t.Run("Expects failed delete categories", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(&categoriesData, nil).Once()
		categoriesRepository.On("DeleteCategories", Id, DeletedBy).Return(business.ErrInternalServerError).Once()
		err := categoiresService.RemoveCategories(Id, AdminId, DeletedBy)
		assert.NotNil(t, err)
	})

	t.Run("Expects failed delete categories, unauthorized access", func(t *testing.T) {
		categoriesRepository.On("GetCategoriesById", Id).Return(&categories.Categories{
			ID:        Id,
			AdminID:   3,
			Parent_id: ParentId,
		}, nil).Once()
		err := categoiresService.RemoveCategories(Id, AdminId, DeletedBy)
		assert.NotNil(t, err)
	})
}

func setup() {

	categoriesSpec = categories.CategoriesSpec{
		AdminID:  AdminId,
		Name:     CategoriesName,
		ParentID: ParentId,
	}

	categoriesData = categories.Categories{
		ID:        Id,
		Name:      CategoriesName,
		Status:    Status,
		Parent_id: ParentId,
		AdminID:   AdminId,
	}

	filterCategories = categories.FilterCategories{}

	categoriesUpdate = categories.CategoriesUpdatable{
		Name:   CategoriesName,
		Status: Status,
	}

	categoiresService = categories.InitCategoriesService(&categoriesRepository)

	listCategories = append(listCategories, categoriesData)
}
