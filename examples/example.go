package main

import (
	"fmt"

	"github.com/focusteach/gfworks/bus"
	"github.com/focusteach/gfworks/examples/routes"
	"github.com/focusteach/gfworks/pkg/conf"
	"github.com/focusteach/gfworks/server/web"
)

var Config web.Conf

func main() {
	conf.Init(true, true)
	err := conf.Load(&Config, "web.yaml")

	fmt.Printf("config: %#v, ret:%v.\n", Config, err)

	app := bus.GetInstance()
	webserver := web.New(Config)

	routes.InitRouter(*webserver)
	app.AddTask(webserver)

	app.Exec()

}
