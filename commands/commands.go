package commands

import (
	"log"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

const triggerTimeUnit = time.Millisecond

type GarageDoorController struct {
	pin              rpio.Pin
	triggered        bool
	triggerTime      time.Duration
	forceTriggerTime time.Duration
	cancel           chan bool
}

func ControllerFactory(pin int, triggerTime int, forceTime int) (controller GarageDoorController) {
	controller = GarageDoorController{
		pin:              rpio.Pin(pin),
		triggered:        false,
		triggerTime:      time.Duration(triggerTime) * triggerTimeUnit,
		forceTriggerTime: time.Duration(forceTime) * triggerTimeUnit,
		cancel:           make(chan bool)}
	controller.pin.Output()
	controller.pin.High()
	return
}

func (controller *GarageDoorController) SetTriggered(triggered bool) {
	controller.triggered = triggered
}

func (controller *GarageDoorController) Trigger(force bool) {
	if controller.triggered {
		controller.cancel <- true
	} else {
		go func() {
			controller.SetTriggered(true)
			controller.pin.Toggle()
			triggerTime := controller.triggerTime
			if force {
				triggerTime = controller.forceTriggerTime
			}

			timeout := time.NewTimer(triggerTime)

			select {
			case <-controller.cancel:
				log.Println("Trigger canceled")
				timeout.Stop()
			case <-timeout.C:
				log.Println("Trigger complete")
			}

			controller.pin.Toggle()
			controller.SetTriggered(false)
		}()
	}
}
