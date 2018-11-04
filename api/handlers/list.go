package handlers

import (
	"net/http"

	"github.com/Sirupsen/logrus"
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
}
