# portcheck 🔌

Check if TCP ports are open on a host.

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/clarabennettdev/portcheck)](https://goreportcard.com/report/github.com/clarabennettdev/portcheck)

Scan individual ports, comma-separated lists, or ranges. Useful for verifying firewall rules and debugging connectivity.

## Install

```bash
go install github.com/clarabennettdev/portcheck@latest
```

## Usage

```bash
portcheck localhost 80,443,8080
portcheck example.com 20-25
portcheck 192.168.1.1 22,80 2000   # custom timeout (ms)
```

## License

MIT
