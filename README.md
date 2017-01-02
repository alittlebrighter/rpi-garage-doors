# Raspberry Pi Garage Door Controller

A simple controller for my garage doors.

Wire up your relay to the pi and your garage door button wires to the relay (http://www.instructables.com/id/Raspberry-Pi-Garage-Door-Opener/step7/Attach-Raspberry-Pi-to-the-Garage/).

Edit config.yml according to the pins you are using (BCM) and change the timings depending on how long the switch needs to be activated to trigger your garage doors.

Assuming you have Go installed on your Pi 
```
$ go install github.com/alittlebrighter/rpi-garage-doors
$ sudo $GOPATH/bin/garage-doors
```
Then make an http POST request to localhost:8080/control with a body of "door=0&force=false".
