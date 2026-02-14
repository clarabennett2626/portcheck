package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func checkPort(host string, port int, timeout time.Duration) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: portcheck <host> <port1,port2,...|port1-port2> [timeout_ms]")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  portcheck localhost 80,443,8080")
		fmt.Println("  portcheck example.com 20-25")
		fmt.Println("  portcheck 192.168.1.1 22,80 2000")
		os.Exit(1)
	}

	host := os.Args[1]
	portSpec := os.Args[2]
	timeout := 1000 * time.Millisecond
	if len(os.Args) > 3 {
		if ms, err := strconv.Atoi(os.Args[3]); err == nil {
			timeout = time.Duration(ms) * time.Millisecond
		}
	}

	var ports []int
	for _, part := range strings.Split(portSpec, ",") {
		part = strings.TrimSpace(part)
		if strings.Contains(part, "-") {
			bounds := strings.SplitN(part, "-", 2)
			start, err1 := strconv.Atoi(bounds[0])
			end, err2 := strconv.Atoi(bounds[1])
			if err1 != nil || err2 != nil || start > end || start < 1 || end > 65535 {
				fmt.Fprintf(os.Stderr, "Invalid range: %s\n", part)
				os.Exit(1)
			}
			for p := start; p <= end; p++ {
				ports = append(ports, p)
			}
		} else {
			p, err := strconv.Atoi(part)
			if err != nil || p < 1 || p > 65535 {
				fmt.Fprintf(os.Stderr, "Invalid port: %s\n", part)
				os.Exit(1)
			}
			ports = append(ports, p)
		}
	}

	fmt.Printf("Scanning %s (%d ports, timeout %v)...\n\n", host, len(ports), timeout)

	type result struct {
		Port int
		Open bool
	}

	results := make([]result, len(ports))
	var wg sync.WaitGroup
	sem := make(chan struct{}, 50) // concurrency limit

	for i, p := range ports {
		wg.Add(1)
		go func(idx, port int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			open := checkPort(host, port, timeout)
			results[idx] = result{port, open}
		}(i, p)
	}
	wg.Wait()

	openCount := 0
	for _, r := range results {
		status := "closed"
		if r.Open {
			status = "OPEN"
			openCount++
		}
		fmt.Printf("  %-6d %s\n", r.Port, status)
	}
	fmt.Printf("\n%d open, %d closed\n", openCount, len(ports)-openCount)
}
