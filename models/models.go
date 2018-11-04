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

// GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	Error string `json:"error"`
}

// PlaceRequest A Pet is the main product in the store.
// It is used to describe the animals available in the store.
//
// swagger:model placeRequest
type PlaceRequest struct {
	// The name of the pet.
	//
	// required: true
	// pattern: \w[\w-]+
	// minimum length: 3
	// maximum length: 50
	Origin []string `json:"origin"`

	// The photo urls for the pet.
	// This only accepts jpeg or png images.
	//
	// required: true
	// items pattern: \.(jpe?g|png)$
	Destination []string `json:"destination"`
}

// PlaceParam request model
//
// This is used for returning a response with a single order as body
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
