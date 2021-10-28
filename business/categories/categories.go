package categories

import (
	"time"
)

type Categories struct {
	ID                     int
	Name                   string
	AdminID                int
	Status                 string
	Parent_id              int
	Count_child_categories int
	Created_at             time.Time
	Updated_at             time.Time
	Deleted_at             time.Time
}

func NewCategories(categories_spec CategoriesSpec) *Categories {
	return &Categories{
		Name:       categories_spec.Name,
		AdminID:    categories_spec.AdminID,
		Status:     "active",
		Parent_id:  categories_spec.ParentID,
		Created_at: time.Now(),
	}
}

func (old_categories *Categories) ModifyOldCategories(categories CategoriesUpdatable) *Categories {
	return &Categories{
		ID:         old_categories.ID,
		Name:       categories.Name,
		AdminID:    old_categories.AdminID,
		Status:     categories.Status,
		Parent_id:  old_categories.Parent_id,
		Created_at: old_categories.Created_at,
		Updated_at: time.Now(),
	}
}
