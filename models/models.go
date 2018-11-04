package models

import (
	"time"
)

//Order order model
type Order struct {
	ID             int64  `json:"id"`
	Distance       int    `json:"distance"`
	Status         string `json:"status"`
	StartLatitude  string
	StartLongitude string
	EndLatitude    string
	EndLongitude   string
	CreateTime     time.Time
	TakenTime      time.Time
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
