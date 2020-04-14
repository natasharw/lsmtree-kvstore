package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetKey(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("GET", "get/", nil)
	if err != nil {
		t.Fatalf("could not create handler: %v", err)
	}
	rr := httptest.NewRecorder()
	http.HandlerFunc(DefaultPage).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status returned. got %v want %v",
			status, http.StatusOK)
	}

	want := "foo" // [TODO]
	if rr.Body.String() != want {
		t.Errorf("wrong body returned. got %v want %v",
			rr.Body.String(), want)
	}
}
