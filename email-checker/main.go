package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	return true
}

func checkDomain(domain string) bool {


	
}