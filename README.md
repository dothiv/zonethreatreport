# Zone Threat Report

This script checks domains in a zone for possible threats.

These resources are used to determine threats:

 - [The Domain Block List](https://www.spamhaus.org/dbl/)

## Usage

    # Fetch zone file
	wget https://tld.hiv/uploads/reports/hiv.zone.gz
	
	# Extract and make uniq
	gunzip -c hiv.zone.gz | awk '{ print $1; }' | sort | uniq > hiv.zone
	
	# Run
	./zonethreatreport hiv.zone hiv 2> zonethreatreport-hiv-`date +%Y-%m-%d`.log

