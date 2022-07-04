# Guardian Tales RPC

<img alt="Rich Presence" src="docs/thumbnail.png">

GT-RPC is a Discord Rich Presence client for Guardian Tales that runs on [rich-go](https://github.com/hugolgst/rich-go).

The activity updates on login and resource changes (such as when coins are collected).

### Components
- Title: Will always be "Guardian Tales".
- Line 1: If inside PvP mode, the mode will be displayed. If resource changes/login action was recorded by the client, the World will be displayed. Otherwise, the client will display "Loading..."
- Line 2: Displays the name and level of the first character of the first team.
- Line 3: If login action was recorded by the client, display the time elapsed since start up. Otherwise, display the time elapsed since the first recorded game state change.

### Usage

Run `go build cmd/rpc.go && ./rpc` to build and run the proxy and Rich Presence client.

If there is no `cert.pem` and `key.pem` present in the folder, Kanterbury will generate a root CA to be installed in your emulator/device, so that HTTPS traffic can be intercepted.

There are configuration options for the CLI which can be accessed via `./rpc --help`.