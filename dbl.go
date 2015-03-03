package zonethreatreport

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type CommandRunner interface {
	Run(string, ...string) ([]byte, error)
}
type RealCommandRunner struct{}

// the real runner for the actual program, actually execs the command
func (r RealCommandRunner) Run(command string, args ...string) ([]byte, error) {
	out, err := exec.Command(command, args...).CombinedOutput()
	return out, err
}

/**
 * Check if domain is listed in Domain Block List
 */
func Dbl(zonefile io.Reader, dbl string, runner CommandRunner, reportlog io.Writer, infolog io.Writer) (alerts int) {
	scanner := bufio.NewScanner(zonefile)
	numchecks := 0
	for scanner.Scan() {
		domain := strings.Split(scanner.Text(), "\t")[0]
		io.WriteString(infolog, fmt.Sprintf("DBL: %s\n", domain))
		dig, err := runner.Run("dig", domain+dbl, "+short")
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
	io.WriteString(reportlog, fmt.Sprintf("Checked domains: %d\n", numchecks))
	io.WriteString(reportlog, fmt.Sprintf("Alerts:          %d\n", alerts))
	return
}
