package main

import (
	"context"
	trfmain "testproject/pkg/rule/TRF_Main"
	util "testproject/pkg/util"
)

import (
	_ "testproject/docs"
)
import (
	"testproject/pkg/middleware"
)
import (
	"flag"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func main() {

	cntxt := util.CustomContext{}
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	trfmain.TRF_Main(&cntxt, config)

	RunProxy()
}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

func run() {
	e := echo.New()
	port := middleware.ConfigEchoNode(e)
	e.Logger.Fatal(e.Start(":" + port))
}
