package test

import (
	"net/http"
	"net/http/httptest"
	"targeting-engine/handler"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	handler.RegisterRoutes(r)
	return r
}

func TestValidDelivery(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/delivery?app=com.gametion.ludokinggame&country=us&os=android", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestEmptyMatch(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/delivery?app=unknown&country=xyz&os=linux", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
}

func TestMissingParams(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/delivery?country=us&os=android", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
