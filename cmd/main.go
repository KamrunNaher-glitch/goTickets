package main

import (
	"gotickets/internal/config"
	"gotickets/internal/server"	
)

func main() {
	// Load Environment Variables
	cfg := config.LoadEnv()
	// connect to the Database
	db := config.ConnectDatabase(cfg)
	// start the Server
	server.Start(db,cfg)
	
}
