package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/GPFAFF/go-lambda/record"
	"github.com/GPFAFF/go-lambda/sqs/send"
)

// BuildReport creates a vehicle data report for Utilization.
func BuildReport(filename string) {
	lines, err := ReadCsv(filename)

	if err != nil {
		panic(err)
	}

	for _, line := range lines[1:] {

		var vehicle record.VehicleData

		if strings.Contains(filename, "active") {
			vehicle = record.CreateActiveVehicleEntry(line)
		} else {
			vehicle = record.CreateTerminatedVehicleEntry(line)
		}

		// push single entry to sqs
		send.Message(vehicle)
	}
}

// ReadCsv accepts a file and returns its content as a multi-dimensional type with lines and each column.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	// Read File into a Variable
	reader := csv.NewReader(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Split entries by the delimiter
	reader.Comma = '|'

	// Read the entire CSV file
	csvData, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return csvData, nil
}
