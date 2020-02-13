package report

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Utilization is the struct which represents vehicle data.
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

// BuildReport creates the report for exporting.
func BuildReport(name string, line []string) {
	var vehicle []Utilization

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
	utilizationJSON, _ := json.Marshal(vehicle)
	fmt.Println("OUTPUT", string(utilizationJSON))
}
