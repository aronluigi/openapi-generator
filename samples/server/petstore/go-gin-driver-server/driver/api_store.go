/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteOrderResponse400 controller response type
type DeleteOrderResponse400 struct {
}

// DeleteOrderResponse404 controller response type
type DeleteOrderResponse404 struct {
}

// DeleteOrderHeaders : request headers
type DeleteOrderHeaders struct {
}

// DeleteOrderQueryParams : query/GET request parameters
type DeleteOrderQueryParams struct {
}

// DeleteOrderPathParams : in path request parameters
type DeleteOrderPathParams struct {
	// ID of the order that needs to be deleted
	OrderId string `json:"orderId" uri:"orderId" binding:"required"`
}

// DeleteOrderFormParams : form/POST request parameters
type DeleteOrderFormParams struct {
}

// DeleteOrderRequest payload object
type DeleteOrderRequest struct {
	Headers *DeleteOrderHeaders
	QueryParams *DeleteOrderQueryParams
	PathParams *DeleteOrderPathParams
	FormParams *DeleteOrderFormParams
}

// Bind : gin bind
func (o *DeleteOrderRequest) Bind (c *gin.Context) error {
    var err error

    o.PathParams = &DeleteOrderPathParams{}
	err = c.ShouldBindUri(o.PathParams)
	if err != nil {
		return err
	}

	return err
}

// NewDeleteOrderHandler constructor
func NewDeleteOrderHandler(cmd Command) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := cmd(c)

		switch resp.(type) {
		case DeleteOrderResponse400:
			render(c, 400, resp)
		case DeleteOrderResponse404:
			render(c, 404, resp)
		default:
			err := fmt.Sprintf("DeleteOrder() command response type mismatch. [%T] != ENUM[DeleteOrderResponse400, DeleteOrderResponse404, ]")
			c.String(http.StatusNotImplemented, err, resp)
		}
	}
}

// GetInventoryResponse200 controller response type
type GetInventoryResponse200 struct {
	Payload map[string]int32
}

// GetInventoryHeaders : request headers
type GetInventoryHeaders struct {
}

// GetInventoryQueryParams : query/GET request parameters
type GetInventoryQueryParams struct {
}

// GetInventoryPathParams : in path request parameters
type GetInventoryPathParams struct {
}

// GetInventoryFormParams : form/POST request parameters
type GetInventoryFormParams struct {
}

// GetInventoryRequest payload object
type GetInventoryRequest struct {
	Headers *GetInventoryHeaders
	QueryParams *GetInventoryQueryParams
	PathParams *GetInventoryPathParams
	FormParams *GetInventoryFormParams
}

// Bind : gin bind
func (o *GetInventoryRequest) Bind (c *gin.Context) error {
    var err error

	return err
}

// NewGetInventoryHandler constructor
func NewGetInventoryHandler(cmd Command) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := cmd(c)

		switch resp.(type) {
		case GetInventoryResponse200:
			render(c, 200, resp)
		default:
			err := fmt.Sprintf("GetInventory() command response type mismatch. [%T] != ENUM[GetInventoryResponse200, ]")
			c.String(http.StatusNotImplemented, err, resp)
		}
	}
}

// GetOrderByIdResponse200 controller response type
type GetOrderByIdResponse200 struct {
	Order
}

// GetOrderByIdResponse400 controller response type
type GetOrderByIdResponse400 struct {
}

// GetOrderByIdResponse404 controller response type
type GetOrderByIdResponse404 struct {
}

// GetOrderByIdHeaders : request headers
type GetOrderByIdHeaders struct {
}

// GetOrderByIdQueryParams : query/GET request parameters
type GetOrderByIdQueryParams struct {
}

// GetOrderByIdPathParams : in path request parameters
type GetOrderByIdPathParams struct {
	// ID of pet that needs to be fetched
	OrderId int64 `json:"orderId" uri:"orderId" binding:"required"`
}

// GetOrderByIdFormParams : form/POST request parameters
type GetOrderByIdFormParams struct {
}

// GetOrderByIdRequest payload object
type GetOrderByIdRequest struct {
	Headers *GetOrderByIdHeaders
	QueryParams *GetOrderByIdQueryParams
	PathParams *GetOrderByIdPathParams
	FormParams *GetOrderByIdFormParams
}

// Bind : gin bind
func (o *GetOrderByIdRequest) Bind (c *gin.Context) error {
    var err error

    o.PathParams = &GetOrderByIdPathParams{}
	err = c.ShouldBindUri(o.PathParams)
	if err != nil {
		return err
	}

	return err
}

// NewGetOrderByIdHandler constructor
func NewGetOrderByIdHandler(cmd Command) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := cmd(c)

		switch resp.(type) {
		case GetOrderByIdResponse200:
			render(c, 200, resp)
		case GetOrderByIdResponse400:
			render(c, 400, resp)
		case GetOrderByIdResponse404:
			render(c, 404, resp)
		default:
			err := fmt.Sprintf("GetOrderById() command response type mismatch. [%T] != ENUM[GetOrderByIdResponse200, GetOrderByIdResponse400, GetOrderByIdResponse404, ]")
			c.String(http.StatusNotImplemented, err, resp)
		}
	}
}

// PlaceOrderResponse200 controller response type
type PlaceOrderResponse200 struct {
	Order
}

// PlaceOrderResponse400 controller response type
type PlaceOrderResponse400 struct {
}

// PlaceOrderHeaders : request headers
type PlaceOrderHeaders struct {
}

// PlaceOrderQueryParams : query/GET request parameters
type PlaceOrderQueryParams struct {
}

// PlaceOrderPathParams : in path request parameters
type PlaceOrderPathParams struct {
}

// PlaceOrderFormParams : form/POST request parameters
type PlaceOrderFormParams struct {
}

// PlaceOrderRequest payload object
type PlaceOrderRequest struct {
	Headers *PlaceOrderHeaders
	QueryParams *PlaceOrderQueryParams
	PathParams *PlaceOrderPathParams
	FormParams *PlaceOrderFormParams
}

// Bind : gin bind
func (o *PlaceOrderRequest) Bind (c *gin.Context) error {
    var err error

	return err
}

// NewPlaceOrderHandler constructor
func NewPlaceOrderHandler(cmd Command) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := cmd(c)

		switch resp.(type) {
		case PlaceOrderResponse200:
			render(c, 200, resp)
		case PlaceOrderResponse400:
			render(c, 400, resp)
		default:
			err := fmt.Sprintf("PlaceOrder() command response type mismatch. [%T] != ENUM[PlaceOrderResponse200, PlaceOrderResponse400, ]")
			c.String(http.StatusNotImplemented, err, resp)
		}
	}
}

