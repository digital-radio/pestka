build:
	go build -o ./dist/amd64/pestka ./src/main.go

build-arm:
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/arm/pestka ./src/main.go

run:
	go run ./src/main.go

style-fix:
	gofmt -w src/

issue-check:
	go vet ./src/...

test:
	go test -v ./tests
