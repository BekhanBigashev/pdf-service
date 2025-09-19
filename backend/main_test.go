package main

import (
	"ilovepdf/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumEndpoint(t *testing.T) {
	r := router.GetRouter()

	req, _ := http.NewRequest("GET", "/sum?first=5&second=7", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"result":12`)
}

func TestSumEndpointInvalidParams(t *testing.T) {
	r := router.GetRouter()

	req, _ := http.NewRequest("GET", "/sum?first=abc&second=5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"params must be int"`)
}
