build:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-w" -o bin/go-wsdl-win64.exe main.go
	env GOOS=windows GOARCH=386 go build -ldflags "-w" -o bin/go-wsdl-win32.exe main.go
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-w" -o bin/go-wsdl-darwin-amd64 main.go
	env GOOS=darwin GOARCH=arm64 go build -ldflags "-w" -o bin/go-wsdl-darwin-arm64 main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags "-w" -o bin/go-wsdl-linux-amd64 main.go
	env GOOS=linux GOARCH=386 go build -ldflags "-w" -o bin/go-wsdl-linux-386 main.go

clean:
	rm -rf bin/*


test:
	go test -v ./...