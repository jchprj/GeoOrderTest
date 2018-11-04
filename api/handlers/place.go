package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jchprj/GeoOrderTest/mgr"
	"github.com/jchprj/GeoOrderTest/models"
)

// PlaceHandler swagger:route POST /orders orders placeHandler
//
// Create an order with an origin and a destination.
//
// Request:
//    default: placeRequest
// Responses:
//    default: genericError
//        200: placeResponse
func PlaceHandler(w http.ResponseWriter, r *http.Request) {
	// logrus.Info("placeHandler")
	if r.Body == nil {
		http.Error(w, "no body", http.StatusBadRequest)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var msg models.PlaceRequest
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpStatus, err := mgr.NewOrder(msg)
	if err != nil {
		http.Error(w, err.Error(), httpStatus)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, (string)(output))
}
