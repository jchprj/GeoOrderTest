package models

import (
	"time"
)

//Order order model
type Order struct {
	OrderID        int64     `xorm:"'order_i_d' pk autoincr bigint(20)"`
	Distance       int       `xorm:"int(11)"`
	Status         string    `xorm:"varchar(200)"`
	StartLatitude  string    `xorm:"varchar(200)"`
	StartLongitude string    `xorm:"varchar(200)"`
	EndLatitude    string    `xorm:"varchar(200)"`
	EndLongitude   string    `xorm:"varchar(200)"`
	CreateTime     time.Time `xorm:"createTime"`
	TakenTime      time.Time `xorm:"takenTime"`
}

// GenericError response model
// GenericError is the default error message that is generated.
//
// swagger:response genericError
type GenericError struct {
	Error string `json:"error"`
}

// PlaceRequest Used when create new order.
// It is used to describe the initial order request.
//
// swagger:model placeRequest
type PlaceRequest struct {
	// Origin latitude and longitude
	// required: true
	Origin []string `json:"origin"`

	// Destination latitude and longitude
	// required: true
	Destination []string `json:"destination"`
}

// PlaceParams request model
//
// This is used for request
//
// swagger:parameters placeHandler
type PlaceParams struct {
	// in: body
	// required: true
	PlaceRequest PlaceRequest
}

// PlaceResponse response model
//
// This is used for returning a response with a single order as body
//
// swagger:response placeResponse
type PlaceResponse struct {
	ID       int64  `json:"id"`
	Distance int    `json:"distance"`
	Status   string `json:"status"`
}

// TakeParams request model
//
// This is used for request
//
// swagger:parameters takeHandler
type TakeParams struct {
	// in: body
	// required: true
	TakeRequest TakeRequest
}

// TakeRequest Used when take an order.
// Status should be "TAKEN"
//
// swagger:model takeRequest
type TakeRequest struct {
	// status should be "TAKEN"
	//
	// required: true
	Status string `json:"status"`
}

// TakeResponse Used as response when take an order.
// Status will be "SUCCESS"
//
// swagger:response takeResponse
type TakeResponse struct {
	// status will be "SUCCESS"
	//
	// required: true
	Status string `json:"status"`
}

// ListParams request model
//
// This is used for request
//
// swagger:parameters listHandler
type ListParams struct {
	// Page index from 1
	// in: path
	// required: true
	Page string `json:"page"`
	// Order numbers per page, between 1 to 50
	// in: path
	// required: true
	Limit string `json:"limit"`
}

// ListResponse response model
//
// This is used for returning a response with orders as body
//
// swagger:response listResponse
type ListResponse struct {
	List []PlaceResponse `json:"list"`
}
