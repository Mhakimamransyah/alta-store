package cart

import (
	"altaStore/business/cart"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type Cart struct {
	ID        uint   `gorm:"id;primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id" validate:"required" gorm:"not null"`
	Status    string `json:"status" validate:"required" gorm:"type:varchar(25); not null"`
	AddressID *uint  `json:"address_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `soft_delete.DeletedAt`
}

type CartDetail struct {
	ID        uint `gorm:"id;primaryKey;autoIncrement"`
	CartID    uint `json:"cart_id" validate:"required" gorm:"not null"`
	ProductID uint `json:"product_id" validate:"required" gorm:"not null"`
	Price     uint `json:"price" validate:"required" gorm:"not null"`
	Quantity  uint `json:"quantity" validate:"required" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func newCartTable(cart cart.Cart) *Cart {

	return &Cart{
		cart.ID,
		cart.UserID,
		cart.Status,
		cart.AddressID,
		cart.CreatedAt,
		cart.UpdatedAt,
		cart.DeletedAt,
	}

}

func newCartDetailTable(cartDetail cart.CartDetail) *CartDetail {

	return &CartDetail{
		cartDetail.ID,
		cartDetail.CartID,
		cartDetail.ProductID,
		cartDetail.Price,
		cartDetail.Quantity,
		cartDetail.CreatedAt,
		cartDetail.UpdatedAt,
		cartDetail.DeletedAt,
	}

}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func NewGormDBRepositoryWithDeleted(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (col *Cart) CartToService() cart.Cart {
	var cart cart.Cart

	cart.ID = col.ID
	cart.UserID = col.UserID
	cart.Status = col.Status
	cart.AddressID = col.AddressID
	cart.CreatedAt = col.CreatedAt
	cart.UpdatedAt = col.UpdatedAt
	cart.DeletedAt = col.DeletedAt

	return cart
}

func (col *CartDetail) CartDetailToService() cart.CartDetail {
	var cartDetail cart.CartDetail

	cartDetail.ID = col.ID
	cartDetail.CartID = col.CartID
	cartDetail.ProductID = col.ProductID
	cartDetail.Price = col.Price
	cartDetail.Quantity = col.Quantity
	cartDetail.CreatedAt = col.CreatedAt
	cartDetail.UpdatedAt = col.UpdatedAt
	cartDetail.DeletedAt = col.DeletedAt

	return cartDetail
}

func (repo *GormRepository) GetActiveCart(UserID uint) (*cart.Cart, error) {
	var cartData Cart

	err := repo.DB.Where("user_id = ?", UserID).Where("status = ?", "active").Where("deleted_at is null").First(&cartData).Error
	if err != nil {
		return nil, err
	}

	activecart := cartData.CartToService()

	return &activecart, nil
}

func (repo *GormRepository) CreateCart(cart cart.Cart) error {

	cartData := newCartTable(cart)
	cartData.ID = 0

	err := repo.DB.Create(cartData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) InsertCartDetail(cartDetail cart.CartDetail) error {
	cartDetailData := newCartDetailTable(cartDetail)
	cartDetailData.ID = 0

	err := repo.DB.Create(cartDetailData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) GetCartDetailByCartID(CartID uint) ([]cart.CartDetail, error) {
	var cartdetails []CartDetail

	err := repo.DB.Where("cart_id = ?", CartID).Where("deleted_at is null").Find(&cartdetails).Error
	if err != nil {
		return nil, err
	}

	var result []cart.CartDetail
	for _, value := range cartdetails {
		result = append(result, value.CartDetailToService())
	}

	return result, nil
}

func (repo *GormRepository) FindProductOnCartDetail(cartID, productID uint) (*cart.CartDetail, error) {

	var cartDetailData CartDetail

	err := repo.DB.Where("cart_id = ?", cartID).Where(" product_id = ?", productID).Where("deleted_at is null").First(&cartDetailData).Error

	if err != nil {
		return nil, err
	}

	cartDetail := cartDetailData.CartDetailToService()

	return &cartDetail, nil

}

func (repo *GormRepository) UpdateQuantity(cartID, productID, qty uint) error {
	if qty != 0 {
		err := repo.DB.Model(&CartDetail{}).Where("cart_id = ?", cartID).Where("product_id = ?", productID).Where("deleted_at is null").Update("quantity", qty).Error
		if err != nil {
			return err
		}
	} else {
		err := repo.DB.Model(&CartDetail{}).Where("cart_id = ?", cartID).Where("product_id = ?", productID).Where("deleted_at is null").Update("deleted_at", time.Now()).Error
		if err != nil {
			return err
		}
	}

	return nil
}
