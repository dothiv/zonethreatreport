package zonethreatreport

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
)

/**
 * Check if domain is listed in Zeus list
 */
func DomainListCheck(domainlist io.Reader, tld string, reportlog io.Writer, infolog io.Writer) (alerts int) {
	scanner := bufio.NewScanner(domainlist)
	numchecks := 0

	var tld_match = regexp.MustCompile("\\." + tld + "$")

	for scanner.Scan() {
		numchecks += 1
		domain := scanner.Text()
		if tld_match.MatchString(domain) {
			io.WriteString(reportlog, fmt.Sprintf("%s is listed!\n", domain))
			alerts += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	io.WriteString(reportlog, fmt.Sprintf("Checked domains: %d\n", numchecks))
	io.WriteString(reportlog, fmt.Sprintf("Alerts:          %d\n", alerts))
	return
}
