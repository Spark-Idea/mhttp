build-macos:
	echo "Compiling for Mac OS"
	GOOS=darwin GOARCH=amd64 	go build -o bin/mac-amd64/mhttp main.go
	GOOS=darwin GOARCH=arm64 	go build -o bin/mac-arm/mhttp main.go
build-linux:
	echo "Compiling for Linux"
	GOOS=linux GOARCH=386 		go build -o bin/linux-386/mhttp main.go
	GOOS=linux GOARCH=amd64 	go build -o bin/linux-amd64/mhttp main.go
	GOOS=linux GOARCH=arm64 	go build -o bin/linux-arm64/mhttp main.go
build-windows:
	echo "Compiling for Windows"
	GOOS=windows GOARCH=386 	go build -o bin/win-386/mhttp.exe main.go
	GOOS=windows GOARCH=amd64 	go build -o bin/win-amd64/mhttp.exe main.go
	GOOS=windows GOARCH=arm64 	go build -o bin/win-arm64/mhttp.exe main.go

clean:
	echo "Removing all contents of bin/ directory"
	rm -rf ./bin/*

build-all: clean build-macos build-linux build-windows
