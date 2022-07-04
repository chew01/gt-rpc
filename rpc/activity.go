package rpc

import (
	"fmt"
	"github.com/chew01/kanterbury/proxy"
	"github.com/chew01/kanterbury/utils"
	"github.com/hugolgst/rich-go/client"
	"time"
)

// ActivityFactory creates a templated instance of client.Activity
func ActivityFactory(gs *proxy.GameState) client.Activity {
	// Convert start time from int type to Time type
	st := time.Unix(0, gs.Startup.StartTime*int64(time.Millisecond))

	var detailString, charString string

	// Priority is given to activity
	if gs.Activity.Name != "" && gs.Activity.EndTime == 0 {
		detailString = gs.Activity.Name
	} else if gs.Player.GameWorld != "" {
		detailString = gs.Player.GameWorld
	} else {
		detailString = ""
	}

	if gs.Character.Name != "" {
		charString = fmt.Sprintf("%s (Lv. %d)", gs.Character.Name, gs.Character.Level)
	} else {
		charString = "Loading..."
	}

	return client.Activity{
		Details:    detailString,
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

// ActivityHook refreshes the RPC activity on each game state update
func ActivityHook(gs *proxy.GameState) {
	act := ActivityFactory(gs)
	utils.Must(client.SetActivity(act))
}
