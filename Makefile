build:
	go build -o ./dist/amd64/pestka ./src/main.go

build-arm:
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/arm/pestka ./src/main.go

run:
	go run ./src/main.go

style-fix:
	gofmt -w src/

style-check:
	echo "FMT check\n"
	gofmt -l ./src/
	echo "\nLINT check\n"
	golint -min_confidence 0 ./src/...

test:
	go test -v ./tests
