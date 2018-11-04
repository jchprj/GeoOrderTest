package handlers_test

import (
	"encoding/json"
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

type request struct {
	start, end    []string
	expectedCode  int
	expectedError string
}

func BenchmarkPlaceHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		postStrings([]string{"12", ""}, []string{"667", ""})
	}
}

func TestPlaceHandler(t *testing.T) {
	cfg.InitConfig("../../config.yml")
	autoID := mgr.GetCurrentAutoID()
	//multiple JSON request test
	tt := []request{
		{
			[]string{"+90.0", "-127.554334"}, []string{"+90.0", "-127.554334"},
			http.StatusOK, `{"id":` + strconv.FormatInt(autoID+1, 10) + `,"distance":0,"status":"UNASSIGNED"}`,
		},
		{
			[]string{"45", "180"}, []string{"+90.0", "-127.554334"},
			http.StatusOK, `{"id":` + strconv.FormatInt(autoID+2, 10) + `,"distance":0,"status":"UNASSIGNED"}`,
		},
		{
			[]string{"+90.0", "-127.554334"}, []string{"+90.0", "-127.554334"},
			http.StatusOK, `{"id":` + strconv.FormatInt(autoID+3, 10) + `,"distance":0,"status":"UNASSIGNED"}`,
		},
		{
			[]string{"+90.0", "-127.554334"}, []string{"+90.0", "-127.554334"},
			http.StatusOK, `{"id":` + strconv.FormatInt(autoID+4, 10) + `,"distance":0,"status":"UNASSIGNED"}`},
		{
			[]string{"heap", "-127.554334"}, []string{"+90.0", "-127.554334"},
			http.StatusBadRequest, `{"error":"ERROR_DESCRIPTION"}`,
		},
		{
			[]string{"+90.0", "-127.554334"}, []string{"", ""},
			http.StatusBadRequest, `{"error":"ERROR_DESCRIPTION"}`,
		},
		{
			[]string{"+90.0"}, []string{""},
			http.StatusBadRequest, `{"error":"ERROR_DESCRIPTION"}`,
		},
	}
	for _, tc := range tt {
		rr, err := postStrings(tc.start, tc.end)
		if err != nil {
			t.Fatal(err)
		}
		checkResponse(rr, tc.expectedCode, tc.expectedError, t)
	}

	//single string test
	rr, err := postString([]byte("abc"))
	if err != nil {
		t.Fatal(err)
	}
	checkResponse(rr, http.StatusBadRequest, `{"error":"INVALID_PARAMETERS"}`, t)
}

func postStrings(start, end []string) (rr *httptest.ResponseRecorder, err error) {
	msg := models.PlaceRequest{
		Origin:      start,
		Destination: end,
	}
	output, _ := json.Marshal(msg)
	return postString(output)
}

func postString(output []byte) (rr *httptest.ResponseRecorder, err error) {
	req, err := http.NewRequest("POST", "/orders", strings.NewReader(string(output)))
	if err != nil {
		return &httptest.ResponseRecorder{}, err
	}
	rr = httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(models.APIPathOrder, handlers.PlaceHandler).Methods("POST")
	router.ServeHTTP(rr, req)
	return rr, nil
}

func checkResponse(rr *httptest.ResponseRecorder, expectedCode int, expected string, t *testing.T) {
	if status := rr.Code; status != expectedCode {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedCode)
	}

	output := strings.Trim(rr.Body.String(), "\n")
	if output != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", output, expected)
	}
}
