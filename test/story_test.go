package test

import (
	"Wizz-Home-Page/Global"
	"Wizz-Home-Page/Init"
	"github.com/gavv/httpexpect/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIndexGetRouter(t *testing.T) {

	if err := os.RemoveAll("./data"); err != nil {
		log.Println(err)
	}

	Init.Init()

	server := httptest.NewServer(Global.Engine)

	e := httpexpect.New(t, server.URL)

	e.GET("/api/stories").
		Expect().
		Status(http.StatusOK).JSON().Array().Empty()

	e.POST("/api/login/auth").WithText(`{
  "password": "admin",
  "username": "admin"
}`).Expect().Status(200)
}
