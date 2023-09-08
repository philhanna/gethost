package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
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

	// Print IP address
	fmt.Printf("addr       = %s\n", addr)

	// Look up the hostname, aliases, and IP addresses
	// associated with the address
	names, err := net.LookupAddr(addr)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		fmt.Printf("name       = %s\n", name)
	}

}

// Handles a host name string
func handleHostName(hostName string) {

	// Print host name
	fmt.Printf("name       = %s\n", hostName)

	// Print canonical name, if available
	cname, err := net.LookupCNAME(hostName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("alias      = %s\n", cname)

	// Look up the IP addresses associated with the hostname
	ips, err := net.LookupIP(hostName)
	if err != nil {
		log.Fatal(err)
	}
	for _, ipAddress := range ips {
		fmt.Printf("addr       = %s\n", ipAddress.String())
	}
}
