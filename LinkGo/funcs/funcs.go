// Utilities and in-and-out operations

package funcs

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// return the first N characters of a string
func FirstN(s string, i int) string {
	return strings.Join(((strings.Split(s, ""))[:i]), "")
}

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
