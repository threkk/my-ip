package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/threkk/myip/pkg/myip"
	"os"
	"strings"
)

// version Version of the software.
var version string

// noColor Disable colouring the output.
var noColor = false

// isPublic If true, retrieves the public addresses.
var isPublic bool

// isLocal If true, retrieves the private addresses.
var isLocal bool

// ipv6 If true, returns IPv6 over IPv4
var ipv6 bool

// isJSON If true, returns the answer in JSON format instead of plain text.
var isJSON bool

// isLong If true, returns the string formatted.
var isLong bool

// isVersion If true, return the version of the software.
var isVersion bool

// isBitbar If true, generate a Bitbar compatible string.
var isBitbar bool

// jsonFormat Type used to define the structure of the JSON response.
type jsonFormat struct {
	Local  []string `json:"local"`
	Public []string `json:"public"`
}

// usage Custom usage function to return in case of error or help message.
func usage(out *os.File, description bool) {
	if description {
		fmt.Fprintf(out, "Returns the IP addresses of the system.\n")
		fmt.Fprintf(out, "\n")
	}
	fmt.Fprintf(out, "Usage: %s [options]\n", os.Args[0])
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "Options:\n")
	flag.PrintDefaults()
}

// bold Turns a string into bold if the noColor flag is not enabled.
func bold(str string) string {
	if noColor {
		return str
	}
	return fmt.Sprintf("\033[1m%s\033[0m", str)
}

// printArr Function used to print an array with address.
func printArr(strs []string, title string) {
	if len(strs) == 0 {
		fmt.Fprintf(os.Stdout, "No addresses were found.\n")
	} else {
		if isLong {
			fmt.Fprintf(os.Stdout, "%s: ", bold(title))
			for _, local := range strs {
				fmt.Fprintf(os.Stdout, "%s ", local)
			}
			fmt.Fprintf(os.Stdout, "\n")
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(strs, " "))
		}
	}
}

func init() {
	if version == "" {
		version = "dev"
	}

	if os.Getenv("NO_COLOR") != "" {
		noColor = true
	}

	flag.BoolVar(&isPublic, "public", false, "Retrieves the public address")
	flag.BoolVar(&isLocal, "local", false, "Retrieves the local addresses")
	flag.BoolVar(&ipv6, "ipv6", false, "Prefer IPv6 over IPv4")
	flag.BoolVar(&isLong, "long", false, "Use long output")
	flag.BoolVar(&isJSON, "json", false, "Export the results in JSON format")
	flag.BoolVar(&isVersion, "version", false, "Display the version number")
	flag.BoolVar(&isBitbar, "bitbar", false, "Generate a Bitbar compatible plugin")

	flag.Usage = func() {
		usage(os.Stdout, true)
	}
}

func main() {
	flag.Parse()

	if isVersion {
		fmt.Fprintf(os.Stdout, "%s\n", version)
		os.Exit(0)
	}

	if isBitbar {
		// TODO
	}

	if !isLocal && !isPublic {
		fmt.Fprintf(os.Stderr, "No type of ip address selected.\n")
		usage(os.Stderr, false)
		os.Exit(1)
	}

	myip.PreferIPv6 = ipv6
	response := &jsonFormat{
		Local:  make([]string, 0),
		Public: make([]string, 0),
	}

	if isLocal {
		locals, err := myip.Local()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Operation failed: %s\n", err)
			os.Exit(1)
		}

		response.Local = locals
		if !isJSON {
			printArr(locals, "Local")
		}
	}

	if isPublic {
		ctxbg := context.Background()
		publics, err := myip.Public(ctxbg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Operation failed: %s\n", err)
			os.Exit(1)
		}

		response.Public = publics
		if !isJSON {
			printArr(publics, "Public")
		}
	}

	if isJSON {
		bytes, err := json.Marshal(response)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Operation failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "%s\n", string(bytes[:]))
		os.Exit(0)
	}
}
