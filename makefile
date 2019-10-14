build:
	/usr/bin/protoc -I. --go_out=plugins=micro:. proto/user.proto