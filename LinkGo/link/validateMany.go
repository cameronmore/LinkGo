package link

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cameronmore/LinkGo/LinkGo/funcs"
)

func ValidateCSV(linkFile, outputFile string, brokenOnly bool) {
	links := funcs.ReadCsvFile(linkFile)
	for _, v := range links {
		// fmt.Println(v[0])
		l := New(v[0])
		err := l.Validate()
		if err != nil {
			fmt.Println("A problem occured while validating", l.URL)
			fmt.Println(err)
			continue
		}
		if !brokenOnly {
			if l.ResponseCode < 400 {
				continue
			}

		}
		// TODO move this to the parent block, out of this loop
		// so the loop doesn't need to open a new file every time?
		// ALSO, this logic cannot comfortably be abstracted away
		// because we need the file's opening and differed closing behavior
		// which is not available with a function as easily?
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
