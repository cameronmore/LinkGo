package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/cameronmore/LinkGo/LinkGo/link"
)

func MainCLI() {
	input := flag.String("i", "NA", "Input CSV file")
	output := flag.String("o", "stdout", "Output JSONL file")
	flag.Parse()

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

	link.ValidateCSV(*input, *output)
}
