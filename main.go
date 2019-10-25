package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fopina/subjack/subjack"
)

var version string = "DEV"
var defaultConfig = "https://raw.githubusercontent.com/fopina/subjack/mymaster/fingerprints.json"

func main() {
	var showVersion bool

	o := subjack.Options{}

	flag.StringVar(&o.Domain, "d", "", "Domain.")
	flag.StringVar(&o.Wordlist, "w", "", "Path to wordlist.")
	flag.IntVar(&o.Threads, "t", 10, "Number of concurrent threads.")
	flag.IntVar(&o.Timeout, "timeout", 10, "Seconds to wait before connection timeout.")
	flag.BoolVar(&o.Ssl, "ssl", false, "Force HTTPS connections (May increase accuracy).")
	flag.BoolVar(&o.All, "a", false, "Find those hidden gems by sending requests to every URL. (default: Requests are only sent to URLs with identified CNAMEs)")
	flag.BoolVar(&o.Verbose, "v", false, "Display more information per each request.")
	flag.StringVar(&o.Output, "o", "", "Output results to file (Subjack will write JSON if file ends with '.json').")
	flag.StringVar(&o.Config, "c", defaultConfig, "Path to configuration file.")
	flag.BoolVar(&o.Manual, "m", false, "Flag the presence of a dead record, but valid CNAME entry.")
	flag.BoolVar(&o.NoColor, "no-color", false, "Disable colored output.")
	flag.BoolVar(&o.IncludeEdge, "e", false, "Include edge takeover cases.")
	flag.BoolVar(&o.Follow, "follow", false, "Follow redirects.")
	flag.BoolVar(&showVersion, "version", false, "Show version.")

	flag.Parse()

	if showVersion {
		fmt.Println("Version: " + version + " (fopina fork)")
		return
	}

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	subjack.Process(&o)
}
