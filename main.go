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

var commonPorts = map[int]string{
	21:   "FTP",
	22:   "SSH",
	23:   "TELNET",
	25:   "SMTP",
	53:   "DNS",
	80:   "HTTP",
	110:  "POP3",
	143:  "IMAP",
	443:  "HTTPS",
	3306: "MySQL",
	5432: "PostgreSQL",
	8080: "HTTP-ALT",
}

func checkPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return
	}
	conn.Close()

	service, exists := commonPorts[port]
	if !exists {
		service = "UNKNOWN"
	}

	fmt.Printf("[+] %d/tcp OPEN (%s)\n", port, service)
}

func parsePortRange(raw string) (int, int, error) {
	parts := strings.Split(raw, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range format")
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil || start > end {
		return 0, 0, fmt.Errorf("invalid numbers")
	}

	return start, end, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: netscanner <host> <start_port-end_port>")
		fmt.Println("example: netscanner scanme.nmap.org 20-100")
		os.Exit(1)
	}

	host := os.Args[1]
	startPort, endPort, err := parsePortRange(os.Args[2])
	if err != nil {
		fmt.Println("[-] Error: Port range must be formatted like 1-1000")
		os.Exit(1)
	}

	fmt.Printf("[*] Scanning %s (ports %d-%d)...\n", host, startPort, endPort)

	var wg sync.WaitGroup
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go checkPort(host, port, &wg)
	}

	wg.Wait()
	fmt.Println("[*] Scan completed.")
}