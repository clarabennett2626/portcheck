# portcheck

Check if TCP ports are open on a host.

## Usage

```
portcheck <host> <port1,port2,...|port1-port2> [timeout_ms]
```

## Examples

```bash
portcheck localhost 80,443,8080
portcheck example.com 20-25
portcheck 192.168.1.1 22,80 2000
```

## Install

```
go install github.com/clarabennett2626/portcheck@latest
```

## License

MIT License
