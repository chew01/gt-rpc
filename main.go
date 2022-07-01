package main

import (
	"github.com/chew01/kanterbury/proxy"
	"github.com/hugolgst/rich-go/client"
	"gt-rpc/rpc"
)

const APP_ID = "992073467514060860"

func main() {

	p := proxy.NewProxy()

	err := client.Login(APP_ID)
	if err != nil {
		panic(err)
	}

	p.State.SetFn(rpc.ActivityHook)

	p.Start()

}
