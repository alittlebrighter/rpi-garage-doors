package commands

import (
	"fmt"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

type GarageDoorController struct {
	pin              rpio.Pin
	lock             bool
	triggerTime      time.Duration
	forceTriggerTime time.Duration
}

func ControllerFactory(pin int, triggerTime int, forceTime int) (controller GarageDoorController) {
	controller = GarageDoorController{pin: rpio.Pin(pin), lock: false, triggerTime: time.Duration(triggerTime), forceTriggerTime: time.Duration(forceTime)}
	controller.pin.Output()
	controller.pin.High()
	return
}

func (controller *GarageDoorController) Lock() bool {

	if controller.lock {
		return false
	} else {
		controller.lock = true
		return true
	}
}

func (controller *GarageDoorController) UnLock() {
	controller.lock = false
}

func (controller *GarageDoorController) Trigger(force bool) {
	if !controller.Lock() {
		return
	}

	controller.pin.Toggle()
	if force {
		time.Sleep(controller.forceTriggerTime * time.Millisecond)
	} else {
		time.Sleep(controller.triggerTime * time.Millisecond)
	}
	controller.pin.Toggle()
	fmt.Println("Garage Door trigger complete!")

	controller.UnLock()
}
