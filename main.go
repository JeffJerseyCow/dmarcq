/*
Copyright © 2024 Chris Powell
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/JeffJerseyCow/dmarcq/dmarc"
	"net"
	"os"
)

func main() {

	var domain string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		domain = scanner.Text()
	}

	dmarcDomain := "_dmarc." + domain

	txtRecords, err := net.LookupTXT(dmarcDomain)
	if err != nil {
		// Silently Ignore
		return
	}

	record, err := dmarc.Analyze(domain, dmarcDomain, txtRecords)

	var vulnerable string
	if record.Status == "invalid" || record.Policy == "none" || record.SubdomainPolicy == "none" ||
		record.Percentage != 100 {
		vulnerable = "vulnerable"
	} else {
		vulnerable = "not vulnerable"
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	} else {
		fmt.Printf("%s,%s,%s,%s,%s,%d,%s,%s\n", record.Domain, record.DMARCDomain, record.Status, record.Policy,
			record.SubdomainPolicy, record.Percentage, record.ReportURIAggregate, vulnerable)
	}
}
