package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"repack/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	utils.InitializeLogger()
}

func TestPostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/packs", postHandler)

	t.Run("Valid Pack Sizes", func(t *testing.T) {
		body := bytes.NewBufferString(`{"packSizes": [1, 2, 3]}`)
		req, _ := http.NewRequest("POST", "/packs", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestOptionsHandler(t *testing.T) {
	router := gin.Default()
	router.OPTIONS("/packs", optionsHandler)

	req, _ := http.NewRequest("OPTIONS", "/packs", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET, POST, OPTIONS", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type", w.Header().Get("Access-Control-Allow-Headers"))
}

func TestGetHandlerValidOrderSize(t *testing.T) {
	router := gin.Default()
	router.GET("/packs", getHandler)

	// init packet size
	packets.Lock()
	packets.Sizes = []int{5, 10, 20}
	packets.Unlock()

	req, _ := http.NewRequest("GET", "/packs?orderSize=25", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"packSizes\":[5,10,20],\"packs\":[1,0,1]}", w.Body.String())
}
