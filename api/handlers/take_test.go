package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jchprj/GeoOrderTest/api/handlers"
	"github.com/jchprj/GeoOrderTest/cfg"
	"github.com/jchprj/GeoOrderTest/mgr"
	"github.com/jchprj/GeoOrderTest/models"
)

//Parallel take same order for many times, count successful times through channel
func TestTakeHandler(t *testing.T) {
	cfg.InitConfig("../../config.yml")
	mgr.Test()
	ID := 3
	fmt.Println("start: ")
	postNum := 21
	takeNum := 10000
	for i := 0; i < postNum; i++ {
		postStrings([]string{"+90.0", "-127.554334"}, []string{"+90.0", "-127.554334"})
	}
	router := mux.NewRouter()
	router.HandleFunc(models.APIPathOrder+"/{orderID}", handlers.TakeHandler).Methods("PATCH")
	c := make(chan int, takeNum)
	for i := 0; i < takeNum; i++ {
		go func() {
			rr, err := takeList(router, strconv.Itoa(ID))
			if err != nil {
				t.Fatal(err)
			}
			t.Log(rr.Code)
			c <- rr.Code
			// checkResponse(rr, http.StatusConflict, `{"error":"ORDER_ALREADY_BEEN_TAKEN"}`, t)
		}()
	}
	ok := 0
	n := 0

	for n < takeNum {
		select {
		case code := <-c:
			if code == http.StatusOK {
				ok++
			}
			n++
		}
	}
	fmt.Println("OK: ", ok, n)
	if ok != 1 {
		t.Errorf("ok is not 1")
	}
	rr, err := takeList(router, strconv.Itoa(ID))
	if err != nil {
		t.Fatal(err)
	}
	checkResponse(rr, http.StatusConflict, `{"error":"ORDER_ALREADY_BEEN_TAKEN"}`, t)
}

func takeList(router *mux.Router, ID string) (rr *httptest.ResponseRecorder, err error) {
	req, err := http.NewRequest("PATCH", "/orders/"+ID, strings.NewReader(`{"status": "TAKEN"}`))
	if err != nil {
		return &httptest.ResponseRecorder{}, err
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr, nil
}
