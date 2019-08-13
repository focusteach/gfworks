package conf

import "github.com/jinzhu/configor"

func init() {

}

var configInstance *configor.Configor

// Init init configor
func Init(debug bool, verbose bool) {
	configInstance = configor.New(&configor.Config{Debug: debug, Verbose: verbose})
}

// Load load configs from fileßßß
func Load(config interface{}, filename string) error {
	if configInstance == nil {
		Init(false, false)
	}
	return configInstance.Load(config, filename)
}
