run:
	go run main.go
build:
	GOOS=windows GOARCH=386 go build main.go -o dist/gohass-mqtt-winclient.exe