package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

// Show usage information

func showUsage() {
	fmt.Println("usage: GetHost <ipaddr>|<name>")
	os.Exit(0)
}

func main() {

	// Check if there are any arguments, and show usage if not

	if len(os.Args) == 1 {
		showUsage()
	}

	// Check whether any of the arguments are help flags, and show usage
	// if so

	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" {
			showUsage()
		}
	}

	// Look up information for each argument

	for _, addr := range os.Args[1:] {

		// Check whether the argument is an IP address or a host name

		if matched, _ := regexp.MatchString(`[\d.]+`, addr); matched {

			// Check whether the IP address is well-formed

			if m, _ := regexp.MatchString(`(\d+)\.(\d+)\.(\d+)\.(\d+)`, addr); !m {
				fmt.Printf("Malformed address: %s\n", addr)
				os.Exit(1)
			}

			// Look up the hostname, aliases, and IP addresses
			// associated with the address

			names, err := net.LookupAddr(addr)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ips, _ := net.LookupIP(addr)

			// Print the hostname, aliases, and IP addresses associated
			// with the address

			PrintNames(names[0], addr, ips)

		} else {

			fmt.Printf("DEBUG: addr=%q", addr)
			// Look up the IP addresses associated with the hostname

			ips, err := net.LookupIP(addr)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Look up the hostname, aliases, and IP addresses
			// associated with the first IP address

			names, err := net.LookupAddr(ips[0].String())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Print the hostname, aliases, and IP addresses associated
			// with the hostname

			PrintNames(names[0], addr, ips)
		}
	}
}

func PrintNames(name string, addr string, ips []net.IP) {
	fmt.Printf("name       = %s\n", name)
	aliases, _ := net.LookupCNAME(addr)
	fmt.Printf("aliases    = %s\n", aliases)
	for _, ip := range ips {
		fmt.Printf("addrs      = %s\n", ip.String())
	}
}