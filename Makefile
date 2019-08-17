default: deps build

deps:
	go get -u github.com/spf13/cobra/cobra
	go get github.com/inconshreveable/mousetrap
	go get github.com/bclicn/color

build:
	go build -o ./bin/samaya-dev

build-all:
	GOOS=linux GOARCH=amd64 go build -o ./bin/samaya-linux64
	GOOS=windows GOARCH=amd64 go build -o ./bin/samaya-win64.exe
	GOOS=darwin GOARCH=amd64 go build -o ./bin/samaya-darwin64

	GOOS=linux GOARCH=386 go build -o ./bin/samaya-linux32
	GOOS=windows GOARCH=386 go build -o ./bin/samaya-win32.exe
	GOOS=darwin GOARCH=386 go build -o ./bin/samaya-darwin32

clean:
	rm -rf ./bin