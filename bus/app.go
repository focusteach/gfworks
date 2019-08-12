package bus

import (
	"context"
	"log"

	// "log"
	"os"
	"os/signal"
	"time"

	"github.com/focusteach/gfworks/app"
)

// Application Application
type Application struct {
	tasks []app.IAppTask
}

// New new a applicaiton
func New() *Application {
	return &Application{}
}

// AddTask add application level task
func (app *Application) AddTask(task app.IAppTask) {
	app.tasks = append(app.tasks, task)

	go func() {
		task.Start()
	}()
}

// Exec exec
func (app *Application) Exec() {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Servers ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, task := range app.tasks {
		if err := task.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
	}

	log.Println("Server exiting")
}
