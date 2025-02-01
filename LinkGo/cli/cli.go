package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/cameronmore/LinkGo/LinkGo/link"
)

func MainCLI() {
	input := flag.String("i", "NA", "Input CSV file")
	output := flag.String("o", "stdout", "Output JSONL file")
	brokenOnly := flag.String("a", "All", "Return all links, or just broken ones")
	flag.Parse()
	broken, err := strconv.ParseBool(*brokenOnly)
	if err != nil {
		broken = false
	}

	if *input == "" {
		fmt.Println("Error: The 'input' flag is required.")
		flag.Usage()
		os.Exit(1)
	}

	if *output == "" {
		fmt.Println("Error: The 'input' flag is required.")
		flag.Usage()
		os.Exit(1)
	}

	link.ValidateCSV(*input, *output, broken)
}
