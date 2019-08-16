deps:
	go get -u github.com/spf13/cobra/cobra

run:
	go run .

build:
	go build -o ./bin/samaya.exe

install:
	cp ./bin/samaya /usr/bin

clean:
	rm -rf ./bin