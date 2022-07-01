package rpc

import (
	"fmt"
	"github.com/chew01/kanterbury/proxy"
	"github.com/chew01/kanterbury/utils"
	"github.com/hugolgst/rich-go/client"
	"time"
)

func ActivityFactory(gs *proxy.GameState) client.Activity {
	st := time.Unix(0, gs.Startup.StartTime*int64(time.Millisecond))
	var worldString, charString string
	if gs.Player.GameWorld != "" {
		worldString = gs.Player.GameWorld
	} else {
		worldString = ""
	}
	if gs.Character.Name != "" {
		charString = fmt.Sprintf("%s (Lv. %d)", gs.Character.Name, gs.Character.Level)
	} else {
		charString = "Loading..."
	}

	return client.Activity{
		Details:    worldString,
		State:      charString,
		LargeImage: "mainlogo",
		LargeText:  "Guardian Tales",
		SmallImage: "",
		SmallText:  "",
		Party:      nil,
		Timestamps: &client.Timestamps{
			Start: &st,
			End:   nil,
		},
		Secrets: nil,
		Buttons: nil,
	}
}

func ActivityHook(gs *proxy.GameState) {
	act := ActivityFactory(gs)
	utils.Must(client.SetActivity(act))
}
