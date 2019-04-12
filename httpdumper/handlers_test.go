package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPDumpHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "/dump", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HTTPRequestDump)

	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
