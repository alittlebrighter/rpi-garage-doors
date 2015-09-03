package util

import (
	yaml "gopkg.in/yaml.v2"

	"io/ioutil"
)

type Config struct {
	Controllers struct {
		Garage_doors struct {
			Trigger_time int
			Force_time   int
			Gpio_pins    struct{ Bcm []int }
		}
	}
	Sockets struct {
		Commands string
	}
}

func ParseConfig(filepath string) (conf Config) {
	config_file, err := ioutil.ReadFile(filepath)
	HandleFatal(err)

	_err := yaml.Unmarshal(config_file, &conf)
	HandleFatal(_err)

	return
}
