package server

import (
	"BIGGO/internal/pkg/storage"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	s := New(":8090", store)
	s.Start()
	return s.newAPI()
}

var s = setupRouter()

func TestHealthEndpoint(t *testing.T) {

	w := httptest.NewRecorder()
	req1, _ := http.NewRequest(http.MethodGet, "/health", nil)
	s.ServeHTTP(w, req1)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestSetGetScalar(t *testing.T) {

	w := httptest.NewRecorder()
	var k Entry = Entry{Value: "ok"}
	z, err := json.Marshal(k)
	if err != nil {
		t.Error(err)
	}
	reader := bytes.NewReader(z)
	req1, _ := http.NewRequest(http.MethodPut, "/scalar/put/test1", reader)
	s.ServeHTTP(w, req1)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())

	w1 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/scalar/get/test1", nil)

	s.ServeHTTP(w1, req2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.Equal(t, string(z), w1.Body.String())

}
