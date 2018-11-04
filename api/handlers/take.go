package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jchprj/GeoOrderTest/mgr"
)

// TakeHandler swagger:route POST /orders orders takeHandler
//
// Take a order, will update order status.
//
// Request:
//    default: takeRequest
// Responses:
//    default: genericError
//        200: takeResponse
func TakeHandler(w http.ResponseWriter, r *http.Request) {
	// logrus.Info("takeHandler")
	vars := mux.Vars(r)
	orderIDStr := vars["orderID"]
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		invalidParameters(w)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		invalidParameters(w)
		return
	}
	var msg map[string]string
	err = json.Unmarshal(b, &msg)
	if err != nil {
		invalidParameters(w)
		return
	}
	status := msg["status"]
	if status != "TAKEN" {
		invalidParameters(w)
		return
	}
	httpStatus, err := mgr.TakeOrder(orderID)
	if err != nil {
		sendError(w, err, httpStatus)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "SUCCESS"}`)
}
