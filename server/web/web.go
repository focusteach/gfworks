package web

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/focusteach/gfworks/pkg/log"
	"github.com/gin-gonic/gin"
)

// Conf server configs
type Conf struct {
	AppName    string `default:"gftest"`
	HttpServer struct {
		Mode          string `default:"release"`
		Addr          string `default:"0.0.0.0:80"`
		MaxUploadSize int64  `default:"8388608"`
	}
}

// Engine httpserver
type Engine struct {
	Router  *gin.Engine
	conf    *Conf
	srv     *http.Server
	logFile *os.File
}

// New new a server
func New(conf Conf) *Engine {
	engine := &Engine{}
	engine.conf = &conf

	gin.SetMode(conf.HttpServer.Mode)

	var err error
	// Logging to a file.
	logFilePath := filepath.Join(log.Dir(), conf.AppName+"-web.log")
	engine.logFile, err = os.Create(logFilePath)
	if err != nil {
		log.Infof("err:%+v", err)
	}
	gin.DefaultWriter = io.MultiWriter(engine.logFile)

	engine.Router = gin.Default()

	engine.Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	engine.Router.Use(gin.Recovery())

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
		log.Infof("listen: %s\n", err)
		return err
	}

	return nil
}

// Shutdown shutdown
func (engine *Engine) Shutdown(ctx context.Context) error {
	defer engine.logFile.Close()
	return engine.srv.Shutdown(ctx)
}
