package main

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
	"github.com/sluggard/myfile/application"
	"github.com/sluggard/myfile/config"
)

func TestReadCustomViaUnmarshaler(t *testing.T) {
	// app := application.NewServer(config.LoadConfig(config.DefaultConfigPath))
	app := application.NewServer(config.InitViperConfig())
	e := httptest.New(t, app.App)
	expectedResponse := `Received: main.config{Addr:"localhost:5678", ServerName:"Iris"}`
	e.OPTIONS("/").WithText("addr: localhost:8080\nserverName: Iris").Expect().
		Status(httptest.StatusOK).Body().Equal(expectedResponse)
}

func TestServiceIpl(t *testing.T) {

}
