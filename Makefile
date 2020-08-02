build:
	go build -o ./dist/amd64/pestka ./src/main.go

build-arm:
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/arm/pestka ./src/main.go

run:
	go run ./src/main.go
