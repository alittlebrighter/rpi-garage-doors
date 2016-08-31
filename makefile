CC=go

build:
	$(CC) build rpi-garage-doors.go

install:
	mv rpi-garage-doors /bin/garage-doors
	cp config.yml /etc/garage-doors.conf
