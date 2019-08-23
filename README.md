# `samaya`
> ( समय​ ) - A time synchronization program that uses HTTP protocol.

<p align="left">
    <a href="https://github.com/vasanthdeveloper/samaya/releases">
        <img src="https://img.shields.io/github/v/release/vasanthdeveloper/samaya?include_prereleases&style=flat-square" alt="">
    </a>
    <a href="https://github.com/vasanthdeveloper/samaya/releases">
        <img src="https://img.shields.io/github/downloads/vasanthdeveloper/samaya/total?style=flat-square" alt="">
    </a>
    <a href="#">
        <img src="https://img.shields.io/github/last-commit/vasanthdeveloper/samaya.svg?style=flat-square" alt="">
    </a>
    <a href="#">
        <img src="https://img.shields.io/github/issues/vasanthdeveloper/samaya.svg?style=flat-square" alt="">
    </a>
    <a href="https://discord.gg/UQuhrxE">
        <img src="https://img.shields.io/discord/600920475341946893?style=flat-square" alt="">
    </a>
    <a href="https://github.com/vasanthdeveloper/samaya/blob/mainline/LICENSE">
        <img src="https://img.shields.io/github/license/vasanthdeveloper/samaya?style=flat-square" alt="">
    </a>
</p>

`samaya` is a command line program for synchronizes your system clock using the HTTPS protocol from <a href="https://worldtimeapi.org" target="_blank">World Time API</a>. `samaya` is very helpful in situations where NTP (Network Time Protocol) packets are blocked or inaccessible.

## Usage
```
____ ____ _  _ ____ _   _ ____ 
[__  |__| |\/| |__|  \_/  |__| 
___] |  | |  | |  |   |   |  |

Usage:
  samaya [flags]

Flags:
  -d, --dry               Fetch the time, but don't update it
  -h, --help              help for samaya
  -t, --timezone string   Set time of that timezone (default "auto")
  -v, --verbose           Show extended output
  -V, --version           Print the version number and exit
```

## Building
#### Dependencies:
| Package | Minimum Version |
|---------|-----------------|
| Go      | v1.12.8         |
| Make    | v4.2.1          |

#### Build Procedure:
```
# Create the namespace
mkdir -p "$(go env GOPATH)/src/vasanthdeveloper.com"

# Clone the project
git clone --single-branch --branch mainline https://github.com/vasanthdeveloper/samaya.git "$(go env GOPATH)/src/vasanthdeveloper.com/samaya"

# Enter the freshly cloned project
cd "$(go env GOPATH)/src/vasanthdeveloper.com/samaya"

# Start the compilation
make

# Install the compiled binaries
sudo make install

# Make samaya sync time, every time we boot the computer
sudo systemctl enable --now samaya
```

## Contributions & Issues
Please feel free to raise an issue for suggestions or bugs. You can also tweet me [@vasanthdevelope](https://twitter.com/vasanthdevelope).

## License
> _The samaya project is released under the [GNU Public License v2](LICENSE). <br> Developed &amp; Maintained By Vasanth Developer. Copyright 2019 © Vasanth Developer._
<hr>

> [vasanth.tech](https://vasanth.tech) &nbsp;&middot;&nbsp;
> YouTube [@vasanthdeveloper](https://youtube.com/vasanthdeveloper?sub_confirmation=1) &nbsp;&middot;&nbsp;
> Twitter [@vasanthdevelope](https://twitter.com/vasanthdevelope) &nbsp;&middot;&nbsp;
> Discord [Vasanth Developer](https://discord.gg/UQuhrxE)