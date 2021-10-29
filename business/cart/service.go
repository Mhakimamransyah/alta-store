package cart

import (
	"altaStore/business"
	"altaStore/util/validator"
	"time"
)

//AddToCartSpec create cart and cart detail spec
type AddToCartSpec struct {
	UserID    uint   `validate:"required"`
	ProductID uint   `validate:"required"`
	Price     uint   `validate:"required"`
	Quantity  uint   `validate:"gte=0"`
	Do        string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) AddToCart(addToCartSpec AddToCartSpec) error {
	err := validator.GetValidator().Struct(addToCartSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	getActiveCart, err := s.repository.GetActiveCart(addToCartSpec.UserID)
	if err != nil {
		cart := NewCart(
			addToCartSpec.UserID,
			"active",
			time.Now(),
		)

		err = s.repository.CreateCart(cart)
		if err != nil {
			return err
		}

		getActiveCart, _ = s.repository.GetActiveCart(addToCartSpec.UserID)
	}

	productOnCart, err := s.repository.FindProductOnCartDetail(getActiveCart.ID, addToCartSpec.ProductID)
	if err == nil {
		if addToCartSpec.Quantity == 0 || (addToCartSpec.Do == "subtraction" && addToCartSpec.Quantity > productOnCart.Quantity) { //delete from cart detail
			err = s.repository.UpdateQuantity(getActiveCart.ID, addToCartSpec.ProductID, 0)
		} else {
			if addToCartSpec.Do == "addition" {
				err = s.repository.UpdateQuantity(getActiveCart.ID, addToCartSpec.ProductID, productOnCart.Quantity+addToCartSpec.Quantity)
			} else {
				err = s.repository.UpdateQuantity(getActiveCart.ID, addToCartSpec.ProductID, productOnCart.Quantity-addToCartSpec.Quantity)
			}
		}

		if err != nil {
			return err
		}
	} else {
		if addToCartSpec.Do == "addition" {
			cartDetail := NewCartDetail(
				getActiveCart.ID,
				addToCartSpec.ProductID,
				addToCartSpec.Price,
				addToCartSpec.Quantity,
				time.Now(),
			)

			err = s.repository.InsertCartDetail(cartDetail)
			if err != nil {
				return err
			}
		} else {
			return business.ErrAddToCart
		}
	}

	return nil
}

func (s *service) GetActiveCart(userID uint) (ActiveCart, error) {
	result := ActiveCart{}
	getActiveCart, err := s.repository.GetActiveCart(userID)
	if err != nil {
		return ActiveCart{}, nil
	}

	cartDetails, err := s.repository.GetCartDetailByCartID(getActiveCart.ID)

	if err != nil {
		result = MergeCart(getActiveCart.ID, getActiveCart.Status, getActiveCart.AddressID, nil)
		return result, nil
	}

	var cartDetailsResult []ActiveCartDetail

	for _, value := range cartDetails {
		cartDetailsResult = append(cartDetailsResult, ToActiveCartDetail(value))
	}

	result = MergeCart(getActiveCart.ID, getActiveCart.Status, getActiveCart.AddressID, cartDetailsResult)

	return result, nil

}
