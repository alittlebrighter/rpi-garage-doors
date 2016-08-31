package util

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Controllers struct {
		Garage_doors struct {
			Trigger_time int
			Force_time   int
			Gpio_pins    struct{ Bcm []int }
		}
	}
	Endpoints struct {
		Host  string
		Paths struct {
			Control string
		}
	}
}

func ParseConfig(filepath string) (conf Config) {
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Cannot read config file at %s\n", filepath)
	}

	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		log.Fatalf("Cannot parse config file at %s\n", filepath)
	}

	return
}
