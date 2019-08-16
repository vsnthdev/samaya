deps:
	go get -u github.com/spf13/cobra/cobra

build:
	go build -o ./bin/samaya

install:
	cp ./bin/samaya /usr/bin

clean:
	rm -rf ./bin