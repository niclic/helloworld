package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "http://localhost/hello", nil)
	w := httptest.NewRecorder()
	Hello(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(string(body), "Hello, World!\n")
	assert.Equal(resp.StatusCode, 200)
}
