package main

import (
	"gitlab.com/jfclockwork/be_clockwork/app"
	"gitlab.com/jfclockwork/be_clockwork/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
