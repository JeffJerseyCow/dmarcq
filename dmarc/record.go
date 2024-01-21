package dmarc

type Record struct {
	Domain             string
	DMARCDomain        string
	Status             string
	Policy             string
	SubdomainPolicy    string
	Percentage         int
	ReportURIAggregate string
}

func New(domain string, dmarcDomain string, status string, policy string, subdomainPolicy string,
	percentage int, reportURIAggregate string) *Record {

	return &Record{
		Domain:             domain,
		DMARCDomain:        dmarcDomain,
		Status:             status,
		Policy:             policy,
		SubdomainPolicy:    subdomainPolicy,
		Percentage:         percentage,
		ReportURIAggregate: reportURIAggregate,
	}
}
