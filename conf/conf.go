package conf

import "github.com/jinzhu/configor"

func init() {

}

// Load load configs from file
func Load(config interface{}, filename string) error {
	return configor.Load(&config, filename)
}
