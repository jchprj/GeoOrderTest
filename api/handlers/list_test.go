package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jchprj/GeoOrderTest/api/handlers"
	"github.com/jchprj/GeoOrderTest/cfg"
	"github.com/jchprj/GeoOrderTest/models"
)

func TestListHandler(t *testing.T) {
	cfg.InitConfig("../../config.yml")
	page := 3
	limit := 10
	n := 21
	for i := 0; i < n; i++ {
		postStrings([]string{"+90.0", "-127.554334"}, []string{"+90.0", "-127.554334"})
	}
	rr, err := getList(strconv.Itoa(page), strconv.Itoa(limit))
	if err != nil {
		t.Fatal(err)
	}
	checkResponse(rr, http.StatusOK, `[{"id":21,"distance":0,"status":"UNASSIGNED"}]`, t)
}

func getList(page, limit string) (rr *httptest.ResponseRecorder, err error) {
	req, err := http.NewRequest("GET", "/orders", nil)
	if err != nil {
		return &httptest.ResponseRecorder{}, err
	}
	q := req.URL.Query()
	q.Add("page", page)
	q.Add("limit", limit)
	req.URL.RawQuery = q.Encode()
	rr = httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(models.APIPathOrder, handlers.ListHandler).Methods("GET")
	router.ServeHTTP(rr, req)
	return rr, nil
}
