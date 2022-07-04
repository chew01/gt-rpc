package main

import (
	"flag"
	"github.com/chew01/kanterbury/proxy"
	"github.com/hugolgst/rich-go/client"
	"gt-rpc/rpc"
	"log"
)

var logPath = flag.String("l", "logs/proxy.log", "file to output log to")
var port = flag.Int("p", 8080, "port for proxy server")

const APP_ID = "992073467514060860"

func main() {
	flag.Parse()
	logFlags := log.Lshortfile | log.Ltime

	options := &proxy.Options{
		LogPath:  *logPath,
		LogFlags: logFlags,
		Port:     *port,
	}

	p := proxy.NewProxy(options)

	err := client.Login(APP_ID)
	if err != nil {
		panic(err)
	}

	err = p.State.AddHook("rpc", rpc.ActivityHook)
	if err != nil {
		panic(err)
	}

	p.Start()
}
