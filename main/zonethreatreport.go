package main

import (
	"fmt"
	ztr "github.com/dothiv/zonethreatreport"
	"log"
	"os"
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
	zonefile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer zonefile.Close()

	// tld := os.Args[2]
	ztr.Dbl(zonefile, "dbl.spamhaus.org", os.Stderr, os.Stdout)
}
