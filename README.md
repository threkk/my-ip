# MyIP
[![GoDoc](https://godoc.org/github.com/threkk/myip?status.svg)](https://godoc.org/github.com/threkk/myip) [![Go Report Card](https://goreportcard.com/badge/github.com/threkk/myip)](https://goreportcard.com/report/github.com/threkk/myip) [![GitHub license](https://img.shields.io/github/license/threkk/myip.svg)](https://github.com/threkk/myip/blob/master/LICENSE.md) [![Twitter](https://img.shields.io/twitter/url/https/github.com/threkk/myip.svg?style=social)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fthrekk%2Fmyip)
> Gets the IP addresses of the system.

## Features
- Works with IPv6.
- Multiple outputs: verbose, short, JSON.
- No third party dependencies (no server request, just DNS look up).
- [Bitbar](https://getbitbar.com/) plugin generator.
- [$NO_COLOR](https://no-color.org) support.

## Install
```
go get -u github.com/threkk/myip/cmd/myip
```

## Usage
Select the networks you want with the options `local` and `public`. If no
network is chosen, it will fail. The output by default is short, use `long` or
`json` to change it.

The option `bitbar` always generates a file in the current directory with the
script to be triggered by Bitbar. This script always asks for the public and
local addresses in short format.

If the global variable `$NO_COLOR` is set, the output will have not formatting.

### Examples
- **Get the local IP in long format.**
```
$ myip -local
Local: 192.168.1.1
```

- **Get both addresses in JSON.**
```
$ myip -local -public -json
{"local":["192.168.1.1"],"public":["1.1.1.1"]}
```

- **Generate a Bitbar plugin.**
```
$ cd /path/to/bitbar/config
$ myip -bitbar
$ ls
myip.60s.sh
```

## Maintainer
Alberto Mtnz ([@threkk](https://threkk.com))

## Contribute
Did you find any problem? Please, check the [issues](https://github.com/threkk/myip/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc)!

## License
BSD-3. See `LICENSE` for more information.
