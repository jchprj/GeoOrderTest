package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/jchprj/GeoOrderTest/mgr"
)

// ListHandler swagger:route POST /orders orders listHandler
//
// List orders.
//
// Responses:
//    default: genericError
//        200: listResponse
func ListHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("listHandler")
	pages := r.URL.Query()["page"]
	if len(pages) == 0 {
		invalidParameters(w)
		return
	}
	page, err := strconv.Atoi(pages[0])
	if err != nil {
		invalidParameters(w)
		return
	}
	if page <= 0 {
		invalidParameters(w)
		return
	}
	limits := r.URL.Query()["limit"]
	if len(limits) == 0 {
		invalidParameters(w)
		return
	}
	limit, err := strconv.Atoi(limits[0])
	if err != nil {
		invalidParameters(w)
		return
	}
	if limit <= 0 || limit > 50 {
		invalidParameters(w)
		return
	}
	list := mgr.GetOrderList(page, limit)
	if len(list) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "[]")
		return
	}
	output, err := json.Marshal(list)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, (string)(output))
}
