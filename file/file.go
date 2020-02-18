package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// VehicleData is a single car report entity.
type VehicleData struct {
	VIN          string `json:"VIN"`
	OrigDealerID string `json:"OrigDealerID"`
	ProgramCode  string `json:"ProgramCode"`
	Date         string `json:"Date"`
	Status       string `json:"Status"`
}

// VehicleReport is a collection of VehicleData.
type VehicleReport struct {
	Vehicles []VehicleData
}

func (data *VehicleReport) addVehicle(item VehicleData) []VehicleData {
	data.Vehicles = append(data.Vehicles, item)
	return data.Vehicles
}

func createTerminatedVehicleReport(line []string) VehicleData {
	return VehicleData{
		VIN:          line[1],
		OrigDealerID: line[2],
		ProgramCode:  line[5],
		Status:       "Terminated",
		Date:         line[9],
	}
}

func createActiveVehicleReport(line []string) VehicleData {
	return VehicleData{
		VIN:          line[1],
		OrigDealerID: line[2],
		ProgramCode:  line[8],
		Status:       "Active",
		Date:         line[9],
	}
}

// BuildReport creates a vehicle data report for Utilization.
func BuildReport(filename string) VehicleReport {
	lines, err := ReadCsv(filename)

	if err != nil {
		panic(err)
	}

	var vr VehicleReport

	fmt.Printf("%T\n", vr)

	for _, line := range lines[1:] {

		var vehicle VehicleData

		if strings.Contains(filename, "active") {
			vehicle = createActiveVehicleReport(line)
		} else {
			vehicle = createTerminatedVehicleReport(line)
		}
		vr.addVehicle(vehicle)
	}

	// resp, err := json.Marshal(vr)
	// fmt.Println("VVV", string(resp))
	// checking the output
	return vr
}

// ReadCsv accepts a file and returns its content as a multi-dimensional type
// with lines and each column.
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

	csvData, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return csvData, nil
}
