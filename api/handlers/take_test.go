package handlers_test

import (
	"encoding/json"
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
	cfg.InitConfig("../../docker/config.yml")
	mgr.Test()
	mgr.InitMgr()
	fmt.Println("start: ")
	postNum := 1
	takeNum := 10
	var resp string
	for i := 0; i < postNum; i++ {
		rr, _ := handlers.PostStrings([]string{"+90.0", "-127.554334"}, []string{"+90.0", "-127.554334"})
		resp = rr.Body.String()
	}
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(resp), &obj)
	if err != nil {
		t.Fatal(err)
	}
	IDFloat64, isOK := obj["id"].(float64)
	if isOK == false {
		t.Errorf("not float64")
	}
	ID := strconv.FormatFloat(IDFloat64, 'f', 0, 64)
	fmt.Println("ID!", ID)
	router := mux.NewRouter()
	router.HandleFunc(models.APIPathOrder+"/{orderID}", handlers.TakeHandler).Methods("PATCH")
	c := make(chan int, takeNum)
	for i := 0; i < takeNum; i++ {
		go func() {
			rr, err := takeOrder(ID)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(rr.Code, rr.Body.String())
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
	rr, err := takeOrder(ID)
	if err != nil {
		t.Fatal(err)
	}
	handlers.CheckResponse(rr, http.StatusConflict, `{"error":"ORDER_ALREADY_BEEN_TAKEN"}`, t)
}

func takeOrder(ID string) (rr *httptest.ResponseRecorder, err error) {
	router := handlers.GetRouter()
	req, err := http.NewRequest("PATCH", "/orders/"+ID, strings.NewReader(`{"status": "TAKEN"}`))
	if err != nil {
		return &httptest.ResponseRecorder{}, err
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr, nil
}
