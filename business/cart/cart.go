package cart

import (
	"time"
)

type Cart struct {
	ID        uint
	UserID    uint
	Status    string
	AddressID *uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type CartDetail struct {
	ID        uint
	CartID    uint
	ProductID uint
	Price     uint
	Quantity  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ActiveCart struct {
	CartID      uint
	Status      string
	AddressID   *uint
	CartDetails []ActiveCartDetail
}

type ActiveCartDetail struct {
	ID        uint
	ProductID uint
	Price     uint
	Quantity  uint
}

//NewCart create new Cart
func NewCart(
	userID uint,
	status string,
	createdAt time.Time) Cart {

	return Cart{
		UserID:    userID,
		Status:    status,
		AddressID: nil,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}

//NewCartDetail create new Cart Detail
func NewCartDetail(
	cartID uint,
	productID uint,
	price uint,
	quantity uint,
	createdAt time.Time) CartDetail {

	return CartDetail{
		CartID:    cartID,
		ProductID: productID,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}

//MergeCart merge ActiveCart and ActiveCartDetail
func MergeCart(
	cartID uint,
	status string,
	addressID *uint,
	cartDetails []ActiveCartDetail) ActiveCart {
	return ActiveCart{
		CartID:      cartID,
		Status:      status,
		AddressID:   addressID,
		CartDetails: cartDetails,
	}
}

//ToActiveCartDetail bind CartDetail struct to ActiveCartDetail struct
func ToActiveCartDetail(cartDetail CartDetail) ActiveCartDetail {
	return ActiveCartDetail{
		ID:        cartDetail.ID,
		ProductID: cartDetail.ProductID,
		Price:     cartDetail.Price,
		Quantity:  cartDetail.Quantity,
	}
}
