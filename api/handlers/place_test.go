package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jchprj/GeoOrderTest/api/handlers"
	"github.com/jchprj/GeoOrderTest/models"
)

func BenchmarkPlaceHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func TestPlaceHandler(t *testing.T) {
	rr, err := test()
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func test() (rr *httptest.ResponseRecorder, err error) {
	msg := models.PlaceRequest{
		Origin:      []string{"12", ""},
		Destination: []string{"667", ""},
	}
	output, _ := json.Marshal(msg)

	req, err := http.NewRequest("POST", "/orders", strings.NewReader(string(output)))
	if err != nil {
		return &httptest.ResponseRecorder{}, err
	}

	rr = httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PlaceHandler)

	handler.ServeHTTP(rr, req)
	return rr, nil
}
