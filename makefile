CC=go

install: clean
	$(CC) build rpi-garage-doors.go
	mv rpi-garage-doors /bin/garage-doors
	cp config.yml /etc/garage-doors.conf
