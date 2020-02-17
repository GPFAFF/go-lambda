package read

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type Utilization struct {
	VIN          string `json:"VIN"`
	OrigDealerID string `json:"OrigDealerID"`
	*ProgramDetails
}

// ProgramDetails represents the vehicle data for active and terminated contracts.
type ProgramDetails struct {
	ProgramCode          string `json:"ProgramCode"`
	ContractDate         string `json:",omitempty"`
	PayoffProcessingDate string `json:",omitempty"`
}

// File reads the proper csv
func File(name string) string {

	csvFile, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot read '%s': %s\n", name, err.Error())
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '|'
	var header = true
	var vehicle []Utilization

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

		if strings.Contains(name, "active") {
			vehicle = append(vehicle, Utilization{
				VIN:          line[1],
				OrigDealerID: line[5],
				ProgramDetails: &ProgramDetails{
					ProgramCode:  line[8],
					ContractDate: line[9],
				},
			})
		} else {
			vehicle = append(vehicle, Utilization{
				VIN:          line[1],
				OrigDealerID: line[2],
				ProgramDetails: &ProgramDetails{
					ProgramCode:          line[5],
					PayoffProcessingDate: line[9],
				},
			})
		}
	}
	utilizationJSON, _ := json.Marshal(vehicle)
	vehicleData := string(utilizationJSON)
	// fmt.Println("OUTPUT", string(utilizationJSON))
	return vehicleData
}
