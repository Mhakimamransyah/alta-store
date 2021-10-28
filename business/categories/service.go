package categories

import (
	"altaStore/business"
	"altaStore/util/validator"
	"errors"
)

type service struct {
	repository Repository
}

type CategoriesSpec struct {
	Name     string `form:"name" validate:"required,max=20"`
	ParentID int    `form:"parent_id"`
	AdminID  int    `form:"admin_id"`
}

type CategoriesUpdatable struct {
	Name   string `form:"name" validate:"required,max=20"`
	Status string `form:"status"`
}

func InitCategoriesService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) FindAllCategories(limit, offset int) (*[]Categories, error) {
	list, err := service.repository.GetCategories(limit, offset)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (service *service) FindAllSubCategories(id_categories, limit, offset int) (*[]Categories, error) {
	list, err := service.repository.GetSubCategories(id_categories, limit, offset)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (service *service) InsertCategories(categories_spec CategoriesSpec, id_admin int, createdBy string) error {
	err := validator.GetValidator().Struct(categories_spec)
	if err != nil {
		return business.ErrInvalidSpec
	}
	// check if categories is not parent
	if categories_spec.ParentID != 0 {
		_, err := service.repository.GetCategoriesById(categories_spec.ParentID)
		if err != nil {
			return errors.New("Parent ID not found")
		}
	}
	categories_spec.AdminID = id_admin
	categories := NewCategories(categories_spec)
	err = service.repository.CreateCategories(categories, createdBy)
	if err != nil {
		return err
	}
	return nil
}

func (service *service) ModifyCategories(categories_updatable CategoriesUpdatable, id_categories int, id_admin int, modifiedBy string) error {
	categories, err := service.repository.GetCategoriesById(id_categories)
	if err != nil {
		return err
	}

	err = validator.GetValidator().Struct(categories_updatable)
	if err != nil {
		return business.ErrInvalidSpec
	}

	// check if admin have authorization to modify category
	if categories.AdminID != id_admin {
		return business.ErrUnauthorized
	}

	new_categories := categories.ModifyOldCategories(categories_updatable)
	err = service.repository.UpdateCategories(new_categories, modifiedBy)
	if err != nil {
		return err
	}
	return nil
}

func (service *service) RemoveCategories(id_categories int, id_admin int, deletedBy string) error {
	categories, err := service.repository.GetCategoriesById(id_categories)
	if err != nil {
		return err
	}

	// check if admin have authorization to modify category
	if categories.AdminID != id_admin {
		return business.ErrUnauthorized
	}

	err = service.repository.DeleteCategories(id_categories, deletedBy)
	if err != nil {
		return err
	}
	return nil
}
