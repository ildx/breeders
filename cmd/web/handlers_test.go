package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
	res := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllDogBreedsHandler)
  handler.ServeHTTP(res, req)

  if res.Code != http.StatusOK {
    t.Errorf("wrong status code; expected 200, got %d", res.Code)
  }
}
