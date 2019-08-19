package conf

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
)

var confPath string

func init() {
	flag.StringVar(&confPath, "conf", "./configs", "default config path")

	if err := os.MkdirAll(confPath, 0777); err != nil {
		log.Fatalf("conf path make failed.err：%+v", err)
	}
}

var configInstance *configor.Configor

// Init init configor
func Init(debug bool, verbose bool) {
	configInstance = configor.New(&configor.Config{Debug: debug, Verbose: verbose})
}

// Dir conf path
func Dir() string {
	return confPath
}

// Load load configs from fileßßß
func Load(config interface{}, filename string) error {
	if configInstance == nil {
		Init(false, false)
	}
	return configInstance.Load(config, filepath.Join(confPath, filename))
}
