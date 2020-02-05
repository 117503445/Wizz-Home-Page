package test

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/Init"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexGetRouter(t *testing.T) {
	Init.Init()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/api/stories", nil)
	Global.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]\n", w.Body.String())
}
