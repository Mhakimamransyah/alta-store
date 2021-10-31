package cart

import (
	"altaStore/business"
	"altaStore/business/products"
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
	repository        Repository
	productRepository products.Repository
}

//NewService Construct user service object
func NewService(
	repository Repository,
	productRepo products.Repository) Service {
	return &service{
		repository,
		productRepo,
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

	var product *products.Products
	product, err = s.productRepository.GetDetailProducts(int(addToCartSpec.ProductID))

	if err != nil {
		return business.ErrProductNotFound
	}

	if addToCartSpec.Do == "addition" {
		if product.Stock < int(addToCartSpec.Quantity) {
			return business.ErrProductOOS
		}
	}

	productOnCart, err := s.repository.FindProductOnCartDetail(getActiveCart.ID, addToCartSpec.ProductID)

	if err == nil {
		operation := "add"
		stock := productOnCart.Quantity
		if addToCartSpec.Quantity == 0 || (addToCartSpec.Do == "subtraction" && addToCartSpec.Quantity > productOnCart.Quantity) { //delete from cart detail
			err = s.repository.UpdateQuantity(getActiveCart.ID, addToCartSpec.ProductID, 0)
			if err != nil {
				return err
			}

			err = s.productRepository.UpdateStocks(int(productOnCart.ProductID), int(stock), operation)
			if err != nil {
				return err
			}
		} else {
			if addToCartSpec.Do == "addition" {
				operation = "min"
				stock = addToCartSpec.Quantity
				err = s.productRepository.UpdateStocks(int(productOnCart.ProductID), int(stock), operation)
				if err != nil {
					return err
				}
				err = s.repository.UpdateQuantity(getActiveCart.ID, addToCartSpec.ProductID, productOnCart.Quantity+addToCartSpec.Quantity)
				if err != nil {
					return err
				}

			} else {
				err = s.repository.UpdateQuantity(getActiveCart.ID, addToCartSpec.ProductID, productOnCart.Quantity-addToCartSpec.Quantity)
				if err != nil {
					return err
				}

				stock = addToCartSpec.Quantity
				err = s.productRepository.UpdateStocks(int(productOnCart.ProductID), int(stock), operation)
				if err != nil {
					return err
				}
			}
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
	var productFromRepo *products.Products
	var simpleProduct ProductCart

	for _, value := range cartDetails {
		productFromRepo, _ = s.productRepository.GetDetailProducts(int(value.ProductID))
		simpleProduct = ToProductCart(*productFromRepo)
		cartDetailsResult = append(cartDetailsResult, ToActiveCartDetail(value, simpleProduct))
	}

	result = MergeCart(getActiveCart.ID, getActiveCart.Status, getActiveCart.AddressID, cartDetailsResult)

	return result, nil

}
