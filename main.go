/*
Copyright © 2024 Chris Powell

*/
package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"github.com/emersion/go-msgauth/dmarc"
)

func analyzeDMARC(domain string, dmarcDomain string, txtRecords []string) (*dmarc.Record, error) {
	for _, record := range txtRecords {

		policy, err := dmarc.Parse(record)
		if err != nil {
			// Silently Ignore
			continue
		}

		return policy, nil
	}

	return nil, fmt.Errorf("Invalid DMARC Record Found")
}

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

	policy, err := analyzeDMARC(domain, dmarcDomain, txtRecords)
	// fmt.Println("Domain,DMARC Domain,Status,Policy,Subdomain Policy,Percent")
	if err != nil {
		fmt.Printf("%s,%s,%s,%s,%s,%s\n", domain, dmarcDomain, "Invalid", "", "", "")
	} else if policy.Percent == nil {
		fmt.Printf("%s,%s,%s,%s,%s,%d\n", domain, dmarcDomain, "Valid", policy.Policy, policy.SubdomainPolicy, 100)
	} else {
		fmt.Printf("%s,%s,%s,%s,%s,%d\n", domain, dmarcDomain, "Valid", policy.Policy, policy.SubdomainPolicy,
			*policy.Percent)
	}
}
