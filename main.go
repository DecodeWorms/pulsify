package main

import (
	"github.com/DecodeWorms/pulsify/config"
	serverutils "github.com/DecodeWorms/pulsify/server-utils"
)

//var c config.Config

func main() {
	_ = config.ImportConfig(config.OSSource{})
	router := serverutils.SetUpRouter()
	serverutils.StartServer(router)
}
