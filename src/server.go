package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	rpio "github.com/stianeikeland/go-rpio"

	"commands"
	"util"
)

const httpPost = "POST"

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	controllers := []commands.GarageDoorController{}

	// takes one command line argument specifying the configuration file path
	conf := util.ParseConfig(os.Args[1])

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

	http.HandleFunc("/garage_door_http", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpPost {
			util.HandleWarning(errors.New("HTTP Method " + r.Method + " not valid for this endpoint."))
			return
		}

		parseFormError := r.ParseForm()
		if parseFormError != nil {
			fmt.Fprintf(w, "Error parsing form data.")
			return
		}

		door, doorErr := strconv.Atoi(r.Form.Get("door"))
		if doorErr != nil || door >= len(controllers) {
			fmt.Fprintf(w, "Error parsing door number.")
			return
		}

		force, forceErr := strconv.ParseBool(r.Form.Get("force"))
		if forceErr != nil {
			fmt.Fprintf(w, "Error parsing force value.")
			return
		}

		go controllers[door].Trigger(force)

		fmt.Fprintf(w, "door: %d, force: %v", door, force)
	})

	log.Fatal(http.ListenAndServe(conf.Sockets.Commands, nil))
}
