#  ___    __________   |  Vasanth Srivatsa (Vasanth Developer)
#  __ |  / /___  __ \  |   -----------------------------------------------
#  __ | / / __  / / /  |  https://github.com/vasanthdeveloper/samaya.git
#  __ |/ /  _  /_/ /   |
#  _____/   /_____/    |  The definition file for "make"
#                      |

# Add automatically the ".exe" at the end of build file
BINEXE := 
ifeq ($(OS),Windows_NT)
	BINEXE += ./bin/samaya-dev.exe
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		BINEXE += ./bin/samaya-dev
	endif
endif

# The first command is the default command which runnns without any extra args and just make
# We get the dependencies and build the executable only for that platform
default: deps build

# Fetch all the required dependencies
deps:
	go get -u github.com/spf13/cobra/cobra
	go get github.com/inconshreveable/mousetrap
	go get github.com/bclicn/color

# Build for current processor architecture and operating system
build:
	go build -o $(BINEXE)

# Build for Linux and Windows
# [NOTE]: This target only works on Linux
build-all:
	GOOS=linux GOARCH=amd64 go build -o ./bin/samaya-linux64
	GOOS=windows GOARCH=amd64 go build -o ./bin/samaya-win64.exe

	GOOS=linux GOARCH=386 go build -o ./bin/samaya-linux32
	GOOS=windows GOARCH=386 go build -o ./bin/samaya-win32.exe

# Clean any generated files
clean:
	rm -rf ./bin