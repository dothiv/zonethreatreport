package main

import (
	"fmt"
	ztr "github.com/dothiv/zonethreatreport"
	"io"
	"log"
	"os"
	"time"
)

func help() {
	io.WriteString(os.Stdout, fmt.Sprintf("Zone Threat Report version %s\n", ztr.VERSION))
	io.WriteString(os.Stdout, fmt.Sprintf("Usage: %s <tld> [list1..n]\n", os.Args[0]))
	io.WriteString(os.Stdout, "  tld         top-level domain (without dot)\n")
	io.WriteString(os.Stdout, "              example: hiv\n")
	io.WriteString(os.Stdout, "  list1..n    space separated list of filenames containing domain names\n")
	io.WriteString(os.Stdout, "              example: zeus.list phistank.list\n")
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	tld := os.Args[1]
	zonefilepath := tld + ".zone"

	start := time.Now()
	io.WriteString(os.Stderr, "##\n")
	io.WriteString(os.Stderr, fmt.Sprintf("# Zone Threat Report version %s\n", ztr.VERSION))
	io.WriteString(os.Stderr, fmt.Sprintf("# Start time: %s\n", start.Format(time.RFC3339)))
	io.WriteString(os.Stderr, fmt.Sprintf("# TLD: %s\n", tld))
	io.WriteString(os.Stderr, fmt.Sprintf("# Zone file: %s\n", zonefilepath))
	for i := 2; i < len(os.Args); i++ {
		io.WriteString(os.Stderr, fmt.Sprintf("# Extra file: %s\n", os.Args[i]))
	}
	io.WriteString(os.Stderr, "##\n\n")

	totalAlerts := 0

	// Domain block list
	dbl := "dbl.spamhaus.org"
	io.WriteString(os.Stderr, fmt.Sprintf("# Check if domains are listed in Domain Block List at %s\n", dbl))
	zonefile, err := os.Open(zonefilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer zonefile.Close()
	totalAlerts += ztr.Dbl(zonefile, dbl, os.Stderr, os.Stdout)
	io.WriteString(os.Stderr, "\n")

	// Extra files
	for i := 2; i < len(os.Args); i++ {
		extrafilepath := os.Args[i]
		io.WriteString(os.Stderr, fmt.Sprintf("# Extra file: %s\n", extrafilepath))
		extrafile, err := os.Open(extrafilepath)
		if err != nil {
			log.Fatal(err)
		}
		totalAlerts += ztr.DomainListCheck(extrafile, tld, os.Stderr, os.Stdout)
		io.WriteString(os.Stderr, "\n")
	}

	// Footer
	end := time.Now()
	io.WriteString(os.Stderr, "##\n")
	io.WriteString(os.Stderr, fmt.Sprintf("# Total alerts: %d\n", totalAlerts))
	io.WriteString(os.Stderr, fmt.Sprintf("# End time:     %s\n", end.Format(time.RFC3339)))
	io.WriteString(os.Stderr, fmt.Sprintf("# Duration:     %s\n", end.Sub(start)))
	io.WriteString(os.Stderr, "##\n")
}
