package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jchprj/GeoOrderTest/models"
)

var router *mux.Router

//GetRouter GetRouter
func GetRouter() *mux.Router {
	if router != nil {
		return router
	}
	r := mux.NewRouter()

	r.HandleFunc(models.APIPathOrder, PlaceHandler).Methods("POST")
	r.HandleFunc(models.APIPathOrder+"/{orderID}", TakeHandler).Methods("PATCH")
	r.HandleFunc(models.APIPathOrder, ListHandler).Methods("GET")
	return r
}

//PostStrings PostStrings for test
func PostStrings(start, end []string) (rr *httptest.ResponseRecorder, err error) {
	msg := models.PlaceRequest{
		Origin:      start,
		Destination: end,
	}
	output, _ := json.Marshal(msg)
	return PostString(output)
}

//PostString PostString for test
func PostString(output []byte) (rr *httptest.ResponseRecorder, err error) {
	router := GetRouter()
	req, err := http.NewRequest("POST", "/orders", strings.NewReader(string(output)))
	if err != nil {
		return &httptest.ResponseRecorder{}, err
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr, nil
}

//CheckResponse CheckResponse for test
func CheckResponse(rr *httptest.ResponseRecorder, expectedCode int, expected string, t *testing.T) {
	if status := rr.Code; status != expectedCode {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedCode)
	}

	output := strings.Trim(rr.Body.String(), "\n")
	if expected != `ignore` && output != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", output, expected)
	}
}
