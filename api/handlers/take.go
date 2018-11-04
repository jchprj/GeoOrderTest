package handlers

import (
	"net/http"

	"github.com/Sirupsen/logrus"
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
	logrus.Info("takeHandler")
}
