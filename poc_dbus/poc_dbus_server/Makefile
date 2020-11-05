build:
	go build -o ./dist/amd64/poc_dbus_server ./server.go

build-arm:
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/arm/poc_dbus_server ./server.go

run:
	go run ./server.go

