package read

import (
	"bufio"
	"encoding/csv"
	"github.com/GPFAFF/go-lambda/report"
	"io"
	"log"
	"os"
)

// File reads the proper csv
func ReadFile(name string) {

	csvFile, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot read '%s': %s\n", name, err.Error())
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '|'
	var header = true

	for {
		line, err := reader.Read()

		// skip header csv values
		if header {
			header = false
			continue
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Cannot read '%s': %s\n", line, err.Error())
		}

		BuildReport(name, line)
	}
}
