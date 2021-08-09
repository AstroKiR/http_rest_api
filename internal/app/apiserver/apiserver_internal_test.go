package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_test1(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test1", nil)
	test1(rec, req)
	assert.Equal(t, rec.Body.String(), "test1")
}

func TestAPIServer_test2(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test2", nil)
	test2(rec, req)
	assert.Equal(t, rec.Body.String(), "test2")
}
