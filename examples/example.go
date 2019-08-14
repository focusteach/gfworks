package main

import (
	"github.com/focusteach/gfworks/bus"
	"github.com/focusteach/gfworks/examples/routes"
	"github.com/focusteach/gfworks/pkg/conf"
	"github.com/focusteach/gfworks/pkg/log"
	"github.com/focusteach/gfworks/server/web"
)

func main() {
	app := bus.GetInstance()

	conf.Init(false, false)
	var config web.Conf
	err := conf.Load(&config, "web.yaml")

	log.Infof("config: %#v, ret:%v.\n", config, err)

	webserver := web.New(config)

	routes.InitRouter(*webserver)
	app.AddTask(webserver)

	app.Exec()

}
