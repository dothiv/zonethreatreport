package zonethreatreport

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

/**
 * Check if domain is listed in Domain Block List
 */
func Dbl(zonefile io.Reader, dbl string, reportlog io.Writer, infolog io.Writer) {
	io.WriteString(reportlog, "++++++++++++++++++++++++++++++++++++\n")
	io.WriteString(reportlog, "+++ Domain Block List checks ... +++\n")
	io.WriteString(reportlog, "++++++++++++++++++++++++++++++++++++\n")
	io.WriteString(reportlog, "\n")
	io.WriteString(reportlog, fmt.Sprintf("Checking with %s\n", dbl))

	scanner := bufio.NewScanner(zonefile)
	numchecks := 0
	alerts := 0
	for scanner.Scan() {
		domain := strings.Split(scanner.Text(), "\t")[0]
		io.WriteString(infolog, fmt.Sprintf("DBL: %s\n", domain))
		cmd := exec.Command("dig", domain+dbl, "+short")
		dig, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		numchecks += 1
		if strings.Contains(string(dig), "127.0.1.") {
			io.WriteString(reportlog, fmt.Sprintf("%s is listed in %s: %s\n", domain, dbl, dig))
			alerts += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	io.WriteString(reportlog, "\n\n")
	io.WriteString(reportlog, fmt.Sprintf("Checked domains: %d\n", numchecks))
	io.WriteString(reportlog, fmt.Sprintf("Alerts:          %d\n", alerts))
	io.WriteString(reportlog, fmt.Sprintf("Clean:           %d\n", numchecks-alerts))
}
