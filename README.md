<h1 align="center">samaya</h1>
<p align="center"><strong>( समय​ ) - A time synchronization program that uses HTTP protocol.</strong></p>
<p align="center">
    <a href="https://github.com/vsnthdev/samaya/releases">
        <img src="https://img.shields.io/github/v/release/vsnthdev/samaya?include_prereleases&style=flat-square">
    </a>
    <a href="https://github.com/vsnthdev/samaya/releases">
        <img src="https://img.shields.io/github/downloads/vsnthdev/samaya/total?style=flat-square">
    </a>
    <a href="https://github.com/vsnthdev/samaya/commits/main">
        <img src="https://img.shields.io/github/last-commit/vsnthdev/samaya.svg?style=flat-square">
    </a>
    <a href="https://github.com/vsnthdev/samaya/issues">
        <img src="https://img.shields.io/github/issues/vsnthdev/samaya.svg?style=flat-square">
    </a>
    <a href="https://github.com/vsnthdev/samaya/blob/main/LICENSE.md">
        <img src="https://img.shields.io/github/license/vsnthdev/samaya?style=flat-square">
    </a>
</p>
<br>

<!-- header -->

`samaya` is a command line program for synchronizes your system clock using the HTTPS protocol from <a href="https://worldtimeapi.org" target="_blank">World Time API</a>. `samaya` is very helpful in situations where NTP (Network Time Protocol) packets are blocked or inaccessible.

## 🚀 Usage
```
samaya
A time synchronization program that uses HTTP protocol.

Usage:
  samaya [flags]

Flags:
  -D, --delay int           Delay the number of seconds before requesting the time
  -d, --dry                 Fetch the time, but don't update it
  -h, --help                help for samaya
  -t, --timezone string     Set time of that timezone (default "auto")
  -v, --verbose             Show extended output
  -V, --version             Print the version number and exit
  -w, --wait for internet   Wait until an active internet connection is established
```

## 🛠 Building
#### Dependencies:
| Package | Recommended Version |
|---------|---------------------|
| Go      | v1.12.8             |
| Make    | v4.2.1              |
| Git     | v2.23.0             |

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

<!-- footer -->

## 📰 License
> The **samaya** project is released under the [MIT license](LICENSE.md). <br> Developed &amp; maintained By Vasanth Srivatsa. Copyright 2021 © Vasanth Developer.
<hr>

> <a href="https://vsnth.dev" target="_blank" rel="noopener">vsnth.dev</a> &nbsp;&middot;&nbsp;
> YouTube <a href="https://vas.cx/videos" target="_blank" rel="noopener">@vasanthdeveloper</a> &nbsp;&middot;&nbsp;
> Twitter <a href="https://vas.cx/twitter" target="_blank" rel="noopener">@vsnthdev</a> &nbsp;&middot;&nbsp;
> Discord <a href="https://vas.cx/discord" target="_blank" rel="noopener">Vasanth Developer</a>