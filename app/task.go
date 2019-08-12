package app

import "context"

// IAppTask application level task
type IAppTask interface {
	Name() string
	Start() error
	Shutdown(context context.Context) error
}
