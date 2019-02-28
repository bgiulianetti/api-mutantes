package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDetectMutantWithEmptyDNA(t *testing.T) {
	server := httptest.NewServer(NewRouter())

	res, err := http.Post(fmt.Sprintf("%s/mutant", server.URL), "application/json", bytes.NewBuffer([]byte("{\"dna\" : []}")))
	if err != nil {
		t.Fatalf("No se pudo realizar el post")
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("El status code debería ser 400, pero se obtuvo %v", res.StatusCode)
	}
	defer server.Close()
}

func TestDetectMutantWithWrongPayload(t *testing.T) {
	server := httptest.NewServer(NewRouter())

	res, err := http.Post(fmt.Sprintf("%s/mutant", server.URL), "application/json", bytes.NewBuffer([]byte("{dna : []}")))
	if err != nil {
		t.Fatalf("No se pudo realizar el post")
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("El status code debería ser 500, pero se obtuvo %v", res.StatusCode)
	}
	defer server.Close()
}
