package common

import (
	"altaStore/business"
	"net/http"
)

type errorBusinessResponseCode string

const (
	errInternalServerError errorBusinessResponseCode = "internal_server_error"
	errHasBeenModified     errorBusinessResponseCode = "data_has_been modified"
	errNotFound            errorBusinessResponseCode = "data_not_found"
	errInvalidSpec         errorBusinessResponseCode = "invalid_spec"
	errLogin               errorBusinessResponseCode = "unauthorize"
	errRegister            errorBusinessResponseCode = "conflict"
	errAddToCart           errorBusinessResponseCode = "invalid_spec"
	errActiveCartNotFound  errorBusinessResponseCode = "cart_not_found"
	errCartDetailEmpty     errorBusinessResponseCode = "active_cart_empty"
	errAddressNotFound     errorBusinessResponseCode = "address_not_found"
	errProductNotFound     errorBusinessResponseCode = "product_not_found"
	errProductOOS          errorBusinessResponseCode = "insufficient_product_stock"
	errTransactionNotFound errorBusinessResponseCode = "transaction_not_found"
	errTransactioAccess    errorBusinessResponseCode = "invalid invoice number"
)

//BusinessResponse default payload response
type BusinessResponse struct {
	Code    errorBusinessResponseCode `json:"code"`
	Message string                    `json:"message"`
	Data    interface{}               `json:"data"`
}

//NewErrorBusinessResponse Response return choosen http status like 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewErrorBusinessResponse(err error) (int, BusinessResponse) {
	return errorMapping(err)
}

//errorMapping error for missing header key with given value
func errorMapping(err error) (int, BusinessResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()
	case business.ErrNotFound:
		return newNotFoundResponse()
	case business.ErrInvalidSpec:
		return newValidationResponse(err.Error())
	case business.ErrHasBeenModified:
		return newHasBeedModifiedResponse()
	case business.ErrLogin:
		return newErrorLogin(err.Error())
	case business.ErrRegister:
		return newErrorRegister(err.Error())
	case business.ErrAddToCart:
		return newErrorRegister(err.Error())
	case business.ErrActiveCartNotFound:
		return newErrorActiveCartNotFound(err.Error())
	case business.ErrCartDetailEmpty:
		return newErrorCartDetailEmpty(err.Error())
	case business.ErrAddressNotFound:
		return newErrorAddressNotFound(err.Error())
	case business.ErrProductNotFound:
		return newErrorProductNotFound(err.Error())
	case business.ErrProductOOS:
		return newErrorProductOOS(err.Error())
	case business.ErrTransactionNotFound:
		return newErrorTransactionNotFound(err.Error())
	case business.ErrTransactionAccess:
		return newErrorTransactionAccess(err.Error())
	}
}

//newInternalServerErrorResponse default internal server error response
func newInternalServerErrorResponse() (int, BusinessResponse) {
	return http.StatusInternalServerError, BusinessResponse{
		errInternalServerError,
		"Internal server error",
		map[string]interface{}{},
	}
}

//newHasBeedModifiedResponse failed to validate request payload
func newHasBeedModifiedResponse() (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errHasBeenModified,
		"Data has been modified",
		map[string]interface{}{},
	}
}

//newNotFoundResponse default not found error response
func newNotFoundResponse() (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errNotFound,
		"Data Not found",
		map[string]interface{}{},
	}
}

//newValidationResponse failed to validate request payload
func newValidationResponse(message string) (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errInvalidSpec,
		"Validation failed " + message,
		map[string]interface{}{},
	}
}

//newErrorLogin failed to Login
func newErrorLogin(message string) (int, BusinessResponse) {
	return http.StatusUnauthorized, BusinessResponse{
		errLogin,
		message,
		map[string]interface{}{},
	}
}

//newErrorRegister failed to Login
func newErrorRegister(message string) (int, BusinessResponse) {
	return http.StatusConflict, BusinessResponse{
		errRegister,
		message,
		map[string]interface{}{},
	}
}

func newErrorAddToCart(message string) (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errAddToCart,
		message,
		map[string]interface{}{},
	}
}

func newErrorActiveCartNotFound(message string) (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errActiveCartNotFound,
		message,
		map[string]interface{}{},
	}
}

func newErrorCartDetailEmpty(message string) (int, BusinessResponse) {
	return http.StatusUnprocessableEntity, BusinessResponse{
		errCartDetailEmpty,
		message,
		map[string]interface{}{},
	}
}

func newErrorAddressNotFound(message string) (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errAddressNotFound,
		message,
		map[string]interface{}{},
	}
}

func newErrorProductNotFound(message string) (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errProductNotFound,
		message,
		map[string]interface{}{},
	}
}

func newErrorProductOOS(message string) (int, BusinessResponse) {
	return http.StatusUnprocessableEntity, BusinessResponse{
		errProductOOS,
		message,
		map[string]interface{}{},
	}
}

func newErrorTransactionNotFound(message string) (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errTransactionNotFound,
		message,
		map[string]interface{}{},
	}
}

func newErrorTransactionAccess(message string) (int, BusinessResponse) {
	return http.StatusForbidden, BusinessResponse{
		errTransactioAccess,
		message,
		map[string]interface{}{},
	}
}
