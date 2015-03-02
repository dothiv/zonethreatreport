package zonethreatreport

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

const DBL = ".dbl.spamhaus.org"

/**
 * Check if domain is listen in Domain Block List
 */
func Dbl(zonefile io.Reader, reportlog io.Writer, infolog io.Writer) {
	io.WriteString(reportlog, "++++++++++++++++++++++++++++++++++++\n")
	io.WriteString(reportlog, "+++ Domain Block List checks ... +++\n")
	io.WriteString(reportlog, "++++++++++++++++++++++++++++++++++++\n")
	io.WriteString(reportlog, "\n")
	io.WriteString(reportlog, fmt.Sprintf("Checking with %s\n", DBL))

	scanner := bufio.NewScanner(zonefile)
	numchecks := 0
	alerts := 0
	for scanner.Scan() {
		domain := strings.Split(scanner.Text(), "\t")[0]
		cmd := exec.Command("dig", domain, "+short")
		dig, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		numchecks += 1
		io.WriteString(infolog, fmt.Sprintf("DBL: %s\n", domain))
		if strings.Contains(string(dig), "127.0.1.") {
			io.WriteString(reportlog, fmt.Sprintf("%s is listed in %s: %s\n", domain, DBL, dig))
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
