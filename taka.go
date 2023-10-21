package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"taka-api/internal/config"
	"taka-api/internal/handler"
	"taka-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/taka-api.yaml", "the config file")

func main() {
	flag.Parse()
	// tidy up this to the config file
	var logConfig logc.LogConf
	logConfig.Level = "debug"
	logc.MustSetup(logConfig)

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//func customizeHttpErrorHandling() {
//	httpx.SetErrorHandler(func(err error) (int, interface{}) {
//		switch _ := err.(type) {
//		// wrap/convert specific error befor sending back
//		//case *errorx.BatchError:
//		//	return http.StatusOK, e.Data()
//		default:
//			return http.StatusInternalServerError, nil
//		}
//	})
//}
