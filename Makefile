run:
	go run main.go
build:
	GOOS=windows GOARCH=386 go build -o dist/gohass-mqtt-winclient.exe