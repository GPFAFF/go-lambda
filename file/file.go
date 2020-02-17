package file

import (
	"encoding/csv"
	"fmt"
	"os"
)

type VehicleReport struct {
	VIN          string `json:"VIN"`
	OrigDealerID string `json:"OrigDealerID"`
	ProgramCode  string `json:"ProgramCode"`
	Date         string `json:"Date"`
	Status       string `json:"Status"`
}

func BuildReport(filename string) {
	lines, err := ReadCsv(filename)

	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := VehicleReport{
			VIN:          line[0],
			OrigDealerID: line[1],
			ProgramCode:  line[2],
			Status:       "Active",
			Date:         line[3],
		}
		fmt.Println(data)
	}
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
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

	reader.Comma = '|'

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return csvData, nil
}
