package handlers

import (
	"encoding/json"
	"errors"
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
		invalidParameters(w)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		invalidParameters(w)
		return
	}
	var msg models.PlaceRequest
	err = json.Unmarshal(b, &msg)
	if err != nil {
		invalidParameters(w)
		return
	}

	resp, httpStatus, err := mgr.NewOrder(msg)
	if err != nil {
		sendError(w, err, httpStatus)
		return
	}

	output, err := json.Marshal(resp)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, (string)(output))
}

func invalidParameters(w http.ResponseWriter) {
	sendError(w, errors.New(models.ErrorInvalidParameters), http.StatusBadRequest)
}

func sendError(w http.ResponseWriter, err error, statusCode int) {
	errResp := models.GenericError{
		Error: err.Error(),
	}
	output, err := json.Marshal(errResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, (string)(output), statusCode)
}
