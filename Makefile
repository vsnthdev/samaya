#  ___    __________   |  Vasanth Srivatsa (Vasanth Developer)
#  __ |  / /___  __ \  |   -----------------------------------------------
#  __ | / / __  / / /  |  https://github.com/vasanthdeveloper/samaya.git
#  __ |/ /  _  /_/ /   |
#  _____/   /_____/    |  The definition file for "make"
#                      |

# Critical Program Information
VERSION := 1.0.1
BUILD := Production
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git rev-parse HEAD)

# Additional Compilation Variables
PREFIX := /usr/local
DESTDIR :=

# Platform dependent automatic program information
ifeq ($(OS),Windows_NT)
	USER := $(shell powershell echo $$env:USERNAME)
	HOSTNAME := $(shell powershell echo $$env:COMPUTERNAME)
	TIME := $(shell powershell Get-Date -Format G)
	BUILDOS := $(shell powershell (Get-WmiObject Win32_OperatingSystem).caption)
	KERNEL := $(shell powershell (Get-CimInstance Win32_OperatingSystem).version)
else
	USER := $(shell echo $$USER)
	HOSTNAME := $(shell echo $$HOSTNAME)
	TIME := $(shell date)
	BUILDOS := $(shell uname -s)
	KERNEL := $(shell uname -r)
endif

# Compilation flags
FLAGS := -ldflags "-X 'main.Version=$(VERSION)' -X 'main.Build=$(BUILD)' -X 'main.GitBranch=$(BRANCH)' -X 'main.CommitHash=$(COMMIT)' -X 'main.Username=$(USER)' -X 'main.Hostname=$(shell hostname)' -X 'main.BuildTime=$(TIME)' -X 'main.BuildOS=$(BUILDOS)' -X 'main.Kernel=$(KERNEL)'"

# Add automatically the ".exe" at the end of build file
BINEXE := 
ifeq ($(OS),Windows_NT)
	BINEXE += ./bin/samaya.exe
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		BINEXE += ./bin/samaya
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
	go get github.com/gosuri/uitable

# Build for current processor architecture and operating system
build:
	go build $(FLAGS) -o $(BINEXE)

# Build for Linux and Windows
# [NOTE]: This target only works on Linux
build-all:
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o ./bin/samaya-linux64
	GOOS=windows GOARCH=amd64 go build $(FLAGS) -o ./bin/samaya-win64.exe

	GOOS=linux GOARCH=386 go build $(FLAGS) -o ./bin/samaya-linux32
	GOOS=windows GOARCH=386 go build $(FLAGS) -o ./bin/samaya-win32.exe

# Install the program and the systemd service
# [NOTE]: This target only works on Linux
install:
	install -Dm644 misc/systemd.service $(DESTDIR)/etc/systemd/system/samaya.service
	install -Dm755 bin/samaya $(DESTDIR)$(PREFIX)/bin/samaya

# Target to create Arch Linux package
archpkg:
	mkdir -p ./cache
	cp ./misc/PKGBUILD ./cache/PKGBUILD
	sed -i -e 's/{{version}}/$(VERSION)/g' ./cache/PKGBUILD
	sed -i -e 's/{{revision}}/1/g' ./cache/PKGBUILD
	cd ./cache && echo "" >> ./PKGBUILD && makepkg -g >> ./PKGBUILD
	cd ./cache && makepkg

# Clean any generated files
clean:
	rm -rf ./bin ./cache
