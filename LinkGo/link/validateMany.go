package link

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cameronmore/LinkGo/LinkGo/funcs"
)

func ValidateCSV(linkFile, outputFile string) {
	links := funcs.ReadCsvFile(linkFile)
	for _, v := range links {
		// fmt.Println(v[0])
		l := New(v[0])
		err := l.Validate()
		if err != nil {
			fmt.Println("A problem occured while validating")
			break
		}
		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		defer f.Close()

		jsonData, err := json.Marshal(l)

		if err != nil {
			return
		}

		n, err := f.Write(jsonData)

		if n, err = f.WriteString("\n"); err != nil {
			fmt.Println(n, err)

		}

	}
}
