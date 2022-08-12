package apiserver_test

import (
	"github.com/adminoid/vuego/internal/apiserver"
	"github.com/adminoid/vuego/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := apiserver.New(config.NewConfig(""))
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.HandleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}