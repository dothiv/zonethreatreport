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
	log.Println(fmt.Sprintf("%s version %s ", os.Args[0], ztr.VERSION))
	log.Println(fmt.Sprintf("Usage: %s zone-file tld", os.Args[0]))
	log.Println("  zone-file:  path to zone file")
	log.Println("  tld:        top-level domain (without dot)")
}

func main() {
	if len(os.Args) != 3 {
		help()
		return
	}
	zonefilepath := os.Args[1]
	zonefile, err := os.Open(zonefilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer zonefile.Close()

	tld := os.Args[2]
	start := time.Now()
	io.WriteString(os.Stderr, "##\n")
	io.WriteString(os.Stderr, fmt.Sprintf("# %s version %s\n", os.Args[0], ztr.VERSION))
	io.WriteString(os.Stderr, fmt.Sprintf("# Start time: %s\n", start.Format(time.RFC3339)))
	io.WriteString(os.Stderr, fmt.Sprintf("# TLD: %s\n", tld))
	io.WriteString(os.Stderr, fmt.Sprintf("# Zone file: %s\n", zonefilepath))
	io.WriteString(os.Stderr, "##\n\n")

	dbl := "dbl.spamhaus.org"
	io.WriteString(os.Stderr, fmt.Sprintf("# Check if domains are listed in Domain Block List at %s\n", dbl))
	// ztr.Dbl(zonefile, dbl, os.Stderr, os.Stdout)
	end := time.Now()
	io.WriteString(os.Stderr, "\n##\n")
	io.WriteString(os.Stderr, fmt.Sprintf("# End time: %s\n", end.Format(time.RFC3339)))
	io.WriteString(os.Stderr, fmt.Sprintf("# Duration: %s\n", end.Sub(start)))
	io.WriteString(os.Stderr, "##\n")
}
