# Zone Threat Report

This script checks domains in a zone for possible threats.

It uses the tlds zone file to query the [Domain Block List](https://www.spamhaus.org/dbl/)]. 
The file is expected to be named `$tld.zone`

All extra arguments are interpreted as list of domain names, those are scanned
for domains of the tld.

The program prints the progress to stdout and the report to stderr.

## Resources

We use these resources determine threats:

 - [The Domain Block List](https://www.spamhaus.org/dbl/)
 - [ZeuS Tracker](https://zeustracker.abuse.ch/)
 - [Malware Domain List](http://www.malwaredomainlist.com/)
 - [Cybercrime Tracker](http://cybercrime-tracker.net/)
 - [VXVault](http://vxvault.siri-urz.net/ViriList.php)
 - [PhishTank](http://www.phishtank.com/)
 - [Joe Wein's List](http://www.joewein.net/)


## Usage

    # This fetches are required lists, runs the check and sends the result via email

    # Fetch zone file
	wget https://tld.hiv/uploads/reports/hiv.zone.gz
	
	# Extract and make uniq
	gunzip -c hiv.zone.gz | awk '{ print $1; }' | sort | uniq > hiv.zone
	
	# Fetch zeus file
	curl 'https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist' > zeus.list
	
	# Fetch malwaredomainlist.com list
	curl http://www.malwaredomainlist.com/hostslist/hosts.txt | awk '{ print $2; }' | sort | uniq > malwaredomain.list 
	
	# Fetch cybercrime-tracker.net list
	curl http://cybercrime-tracker.net/all.php | sed "s/<br \\/>/\n/g" | sed -r "s/^([a-z0-9\\.-]+).*/\1/g" | sort | uniq > cybercrime-tracker.list
	
	# Fetch VX Vault list
	curl 'http://vxvault.siri-urz.net/URL_List.php' | sed -r "s/^https*:\\/\\/([a-z0-9\\.-]+).*/\1/g" | sort | uniq > vxvault.list
	
	# Fetch phistank list
	curl http://data.phishtank.com/data/online-valid.csv | awk -F , '{ print $2 }' | sed -r "s/^https*:\\/\\/([a-z0-9\\.-]+).*/\1/g" | sort | uniq > phistank.list
	
	# Joe Wein's list
	curl http://www.joewein.net/dl/bl/dom-bl.txt | awk -F , '{ print $1 }' | sed -r "s/^https*:\\/\\/([a-z0-9\\.-]+).*/\1/g" > joewein.list
	
	# Run
	./zonethreatreport hiv `echo *.list` 2> zonethreatreport-hiv-`date +%Y-%m-%d`.log
	
	# Send report as email
	echo "" | mail -s "TLD threats report `date +%Y-%m-%d`" -a zonethreatreport-hiv-`date +%Y-%m-%d`.log example@nic.tld
