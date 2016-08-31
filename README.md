# Raspberry Pi Garage Door Controller

A simple controller for my garage doors.

Edit config.yml according to the pins you are using (BCM) and change the timings depending on how long the switch needs to be activated to trigger your garage doors.

Assuming you have Go and make installed on your Pi start the project clone the project then
```
$ go get github.com/alittlebrighter/rpi-garage-doors
$ cd $GOPATH/src/github.com/alittlebrighter/rpi-garage-doors
$ make
$ sudo make install
$ sudo garage-doors
```
Then make an http POST request to localhost:8080/control with a body of "door=0&force=false".