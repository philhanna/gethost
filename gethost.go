package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	// Get the IP address or host name
	addr := flag.Arg(0)
	re := regexp.MustCompile(`(\d+).(\d+).(\d+).(\d+)`)
	if re.MatchString(addr) {
		handleIPAddress(addr)
	} else {
		handleHostName(addr)
	}
}

// Handles an IP address of the form 999.999.999.999
func handleIPAddress(addr string) {

	lines := make([]string, 0)
	
	lines = append(lines, "addr," + addr)

	// Look up the hostname, aliases, and IP addresses
	// associated with the address
	names, err := net.LookupAddr(addr)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		lines = append(lines, "name," + name)
	}

	printLines(lines)
}

// Handles a host name string
func handleHostName(hostName string) {

	lines := make([]string, 0)

	// Print host name
	lines = append(lines, "name," + hostName)

	// Print canonical name, if available
	cname, err := net.LookupCNAME(hostName)
	if err != nil {
		log.Fatal(err)
	}
	lines = append(lines, "canonical name," + cname)

	// Look up the IP addresses associated with the hostname
	ips, err := net.LookupIP(hostName)
	if err != nil {
		log.Fatal(err)
	}
	for _, ipAddress := range ips {
		lines = append(lines, "addr," + ipAddress.String())
	}

	printLines(lines)
}

func printLines(lines []string) {
	keys := make([]string, 0)
	values := make([]string, 0)
	maxWidth := 0
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		key, value := tokens[0], tokens[1]
		width := len(key)
		if width > maxWidth {
			maxWidth = width
		}
		keys = append(keys, key)
		values = append(values, value)
	}

	for i := 0; i < len(keys); i++ {
		key := keys[i] + ":"
		for len(key) < maxWidth+4 {
			key += " "
		}
		value := values[i]
		fmt.Printf("%s %s\n", key, value)
	}
}