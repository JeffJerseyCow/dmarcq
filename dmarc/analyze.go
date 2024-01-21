package dmarc

import (
	"fmt"
	msgauthdmarc "github.com/emersion/go-msgauth/dmarc"
)

func Analyze(domain string, dmarcDomain string, txtRecords []string) (*Record, error) {
	for _, txtRecord := range txtRecords {

		rawDmarcRecord, err := msgauthdmarc.Parse(txtRecord)
		if err != nil {
			// Silently Ignore
			continue
		}

		status := "valid"
		policy := fmt.Sprintf("%s", rawDmarcRecord.Policy)

		var subdomainPolicy string
		if rawDmarcRecord.SubdomainPolicy != "" {
			subdomainPolicy = fmt.Sprintf("%s", rawDmarcRecord.SubdomainPolicy)
		} else {
			subdomainPolicy = fmt.Sprintf("%s", rawDmarcRecord.Policy)
		}

		var percent int
		if rawDmarcRecord.Percent != nil {
			percent = *rawDmarcRecord.Percent
		} else {
			percent = 100
		}

		var reportURIAggregate string
		if rawDmarcRecord.ReportURIAggregate != nil {
			reportURIAggregate = fmt.Sprintf("%s", rawDmarcRecord.ReportURIAggregate)
		} else {
			reportURIAggregate = "nil"
		}

		return New(domain, dmarcDomain, status, policy, subdomainPolicy, percent, reportURIAggregate), nil
	}

	return New(domain, dmarcDomain, "invalid", "nil", "nil", 0, "nil"), nil
}
