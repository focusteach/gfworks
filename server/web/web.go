package web

import (
	"context"
	"net/http"
	"time"
	"github.com/focusteach/gfworks/pkg/logmgr"
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
	logName string
}

func accessLog() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// 这里是指定日志打印出来的格式。分别是状态码，执行时间,请求ip,请求方法,请求路由(等下我会截图)
		logmgr.Logf("access",logmgr.InfoLevel, "| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

// New new a server
func New(conf Conf) *Engine {
	engine := &Engine{}
	engine.conf = &conf

	gin.SetMode(conf.HttpServer.Mode)

	engine.Router = gin.Default()

	// 日志打印
	engine.Router.Use(accessLog())

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
		logmgr.Logf("", logmgr.InfoLevel, "listen: %s\n", err)
		return err
	}

	return nil
}

// Shutdown shutdown
func (engine *Engine) Shutdown(ctx context.Context) error {
	return engine.srv.Shutdown(ctx)
}
