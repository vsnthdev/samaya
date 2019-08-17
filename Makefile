default: deps build

deps:
	go get -u github.com/spf13/cobra/cobra
	go get github.com/inconshreveable/mousetrap

run:
	go run .

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/samaya-linux64
	GOOS=windows GOARCH=amd64 go build -o ./bin/samaya-win64.exe
	GOOS=darwin GOARCH=amd64 go build -o ./bin/samaya-darwin64

	GOOS=linux GOARCH=386 go build -o ./bin/samaya-linux32
	GOOS=windows GOARCH=386 go build -o ./bin/samaya-win32.exe
	GOOS=darwin GOARCH=386 go build -o ./bin/samaya-darwin32

install:
	cp ./bin/samaya /usr/bin

clean:
	rm -rf ./bin