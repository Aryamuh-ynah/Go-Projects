package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord, hasDKIM, isValid\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	
}

func checkDomain(domain string) bool {

	var hasMX, hasSPF, hasDMARC, hasDKIM, isValid bool

	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error looking up MX records for domain %s: %v", domain, err)
	} 
	if (len(mxRecords) > 0) {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for domain %s: %v", domain, err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for domain %s: %v", domain, err)
	} 

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%s, %t, %t, %s, %t, %s, %t, %t\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord, hasDKIM, isValid)

	return hasMX && hasSPF && hasDMARC

}