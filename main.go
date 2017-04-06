package main

import (
	"flag"
	"os"

	"github.com/golang-cloud/golang-cloud-config/server"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port    string
	gitURI  string
	basedir = "/tmp/fae-scc"
)

func init() {
	flag.StringVar(&port, "port", ":8888", "port")
	flag.StringVar(&gitURI, "git", "", "server config center git URL")
}

func main() {
	flag.Parse()

	if gitURI == "" {
		flag.Usage()
		os.Exit(1)
	}

	os.RemoveAll(basedir)

	scc := echo.New()

	scc.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[INFO] method=${method}, uri=${uri}, status=${status}\n",
	}))

	hander := server.NewHander(gitURI, basedir)

	scc.GET("/:name/:profiles/:label", hander.Labelled)
	scc.Logger.Fatal(scc.Start(port))
}
