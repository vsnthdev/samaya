deps:
	echo "Will add the dependencies here"

build:
	go build -o ./bin/samaya main.go

install: build
	cp ./bin/samaya /usr/bin

clean:
	rm -rf ./bin