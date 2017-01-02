package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	rpio "github.com/stianeikeland/go-rpio"

	"github.com/alittlebrighter/rpi-garage-doors/commands"
	"github.com/alittlebrighter/rpi-garage-doors/util"
)

const (
	httpGet  = "GET"
	httpPost = "POST"
)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer rpio.Close()

	controllers := []commands.GarageDoorController{}

	// takes one command line argument specifying the configuration file path
	configFile := "/etc/garage-doors.conf"
	if len(os.Args) == 2 {
		configFile = os.Args[1]
	}
	conf := util.ParseConfig(configFile)

	for _, door := range conf.Controllers.Garage_doors.Gpio_pins.Bcm {
		controllers = append(controllers,
			commands.ControllerFactory(door, conf.Controllers.Garage_doors.Trigger_time, conf.Controllers.Garage_doors.Force_time))
	}

	cpuCount := runtime.NumCPU()
	if cpuCount < len(controllers) {
		runtime.GOMAXPROCS(cpuCount)
	} else {
		runtime.GOMAXPROCS(len(controllers))
	}

	http.HandleFunc(conf.Endpoints.Paths.Control, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpPost {
			fmt.Fprintf(w, "HTTP Method "+r.Method+" not valid for this endpoint.")
			return
		}

		door, doorErr := strconv.Atoi(r.PostFormValue("door"))
		if doorErr != nil || door >= len(controllers) {
			fmt.Fprintf(w, "Error parsing door number.")
			return
		}

		force, forceErr := strconv.ParseBool(r.PostFormValue("force"))
		if forceErr != nil {
			force = false
		}

		log.Println("Opening door " + r.PostFormValue("door") + ", force: " + strconv.FormatBool(force))

		go controllers[door].Trigger(force)

		fmt.Fprintf(w, "door: %d, force: %v", door, force)
	})

	log.Println("Starting server at " + conf.Endpoints.Host)
	log.Fatal(http.ListenAndServe(conf.Endpoints.Host, nil))
}
