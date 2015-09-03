# Raspberry Pi Garage Door Controller

A simple controller for my garage doors.

Edit config.yml according to the pins you are using (BCM) and change the timings depending on how long the switch needs to be activated to trigger your garage doors.

Assuming you have Go and make installed on your Pi start the project clone the project then
```
$ cd rpi-garage-doors
$ make deps
$ make build
$ ./run.sh
```
