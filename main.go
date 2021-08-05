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
	"log"
	"net"
	"testproject/pkg/controller"
	"testproject/pkg/xiLogger"
	"testproject/pub"
)
import (
	"flag"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func main() {

	controller.InitGormCrypto() // GormCrypto

	cntxt := util.CustomContext{}
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	trfmain.TRF_Main(&cntxt, config)

	RunProxy()
	RunServer()
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

func RunServer() {
	port := middleware.GetGrpcPort()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := middleware.GetGrpcServerConfigs()

	xiLogger.Log.Info("\n Server listening on port %v \n", ":"+port)
	pub.RegisterServiceServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
