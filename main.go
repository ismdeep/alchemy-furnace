package main

import (
	"github.com/ismdeep/alchemy-furnace/api"
	_ "github.com/ismdeep/alchemy-furnace/cron"
)

func main() {
	api.Run()
}
