package main

import (
	"fmt"

	"github.com/focusteach/gfworks/bus"
	"github.com/focusteach/gfworks/conf"
	"github.com/focusteach/gfworks/server/web"
)

var Config web.Conf

func main() {
	err := conf.Load(&Config, "config.toml")
	fmt.Printf("config: %#v, ret:%v.\n", Config, err)
	app := bus.New()
	webserver := web.New(Config)

	initRoutes(*webserver)
	app.AddTask(webserver)

	app.Exec()

}
