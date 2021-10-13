package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	"github.com/guogeer/quasar/cmd"
	"github.com/guogeer/quasar/log"
	"github.com/guogeer/quasar/util"
)

var id = flag.String("id", "w_gateway", "server id")
var port = flag.Int("port", 8201, "gateway server port")
var proxy = flag.String("proxy", "", "gateway server proxy addr")

func main() {
	flag.Parse()

	log.Infof("start gateway, listen %d", *port)
	addr := fmt.Sprintf("%s:%d", *proxy, *port)
	cmd.RegisterService(&cmd.ServiceConfig{Id: *id, Name: "gateway", Addr: addr})
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Error(err)
			log.Errorf("%s", buf)
		}
	}()

	for {
		util.GetTimerSet().RunOnce()
		// handle message
		cmd.RunOnce()
	}
}
