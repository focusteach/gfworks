package web

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Conf server configs
type Conf struct {
	AppName    string `default:"app name"`
	HttpServer struct {
		Addr          string `default:"0.0.0.0:80"`
		MaxUploadSize int64  `default:"8388608"`
	}
}

// Engine httpserver
type Engine struct {
	Router *gin.Engine
	conf   *Conf
	srv    *http.Server
}

// New new a server
func New(conf Conf) *Engine {
	engine := &Engine{}
	engine.conf = &conf

	engine.Router = gin.Default()

	engine.Router.MaxMultipartMemory = conf.HttpServer.MaxUploadSize // 8 MiB

	return engine
}

// Name app name
func (engine *Engine) Name() string {
	return engine.conf.AppName
}

// Start listen and serve bm engine by given DSN.
func (engine *Engine) Start() error {
	engine.srv = &http.Server{
		Addr:    engine.conf.HttpServer.Addr,
		Handler: engine.Router,
	}

	// service connections
	if err := engine.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
		return err
	}

	return nil
}

// Shutdown shutdown
func (engine *Engine) Shutdown(ctx context.Context) error {
	return engine.srv.Shutdown(ctx)
}
