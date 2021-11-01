package transaction

import (
	"altaStore/api/common"
	"altaStore/api/v1/transaction/request"
	"altaStore/business"
	"altaStore/business/transaction"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service transaction.Service
}

//NewController Construct item API controller
func NewController(service transaction.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) Checkout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	checkoutRequest := new(request.CheckoutRequest)
	checkoutRequest.UserID = uint(userID)

	if err := c.Bind(checkoutRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	response, err := controller.service.Checkout(*checkoutRequest.ToCheckoutSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(response))

}

func (controller *Controller) GetAllTransaction(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	response, err := controller.service.GetAllTransaction(uint(userID))
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) GetTransactionByInvoice(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	invoice := c.QueryParam("invoiceNumber")
	println("---------------------------------")
	println(invoice)
	println("---------------------------------")
	if invoice == "" {
		return c.JSON(common.NewErrorBusinessResponse(business.ErrInvalidSpec))
	}
	response, err := controller.service.FindTransactionByInvoice(invoice, uint(userID))
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(response))
}
