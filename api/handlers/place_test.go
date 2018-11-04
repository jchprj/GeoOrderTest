package handlers_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jchprj/GeoOrderTest/api/handlers"
	"github.com/jchprj/GeoOrderTest/cfg"
	"github.com/jchprj/GeoOrderTest/mgr"
)

type request struct {
	start, end    []string
	expectedCode  int
	expectedError string
}

func BenchmarkPlaceHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		handlers.PostStrings([]string{"12", ""}, []string{"667", ""})
	}
}

func TestPlaceHandler(t *testing.T) {
	cfg.InitConfig("../../config.yml")
	mgr.InitMgr()
	autoID := mgr.GetCurrentAutoID()
	t.Logf("start autoID: %v", autoID)
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
		rr, err := handlers.PostStrings(tc.start, tc.end)
		if err != nil {
			t.Fatal(err)
		}
		handlers.CheckResponse(rr, tc.expectedCode, tc.expectedError, t)
	}

	//single string test
	rr, err := handlers.PostString([]byte("abc"))
	if err != nil {
		t.Fatal(err)
	}
	handlers.CheckResponse(rr, http.StatusBadRequest, `{"error":"INVALID_PARAMETERS"}`, t)
}
