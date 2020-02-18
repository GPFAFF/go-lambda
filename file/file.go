package file

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// VehicleReport is a collection of VehicleData.
type VehicleReport struct {
	VehicleData []VehicleData
}

// VehicleData is a single car report entity.
type VehicleData struct {
	VIN          string `json:"VIN"`
	OrigDealerID string `json:"OrigDealerID"`
	ProgramCode  string `json:"ProgramCode"`
	Date         string `json:"Date"`
	Status       string `json:"Status"`
}

func (v *VehicleReport) addVehicle(item VehicleData) []VehicleData {
	v.VehicleData = append(v.VehicleData, item)
	return v.VehicleData
}

func createTerminatedVehicleReport(line []string) VehicleData {
	currentVehicle := VehicleData{
		VIN:          line[1],
		OrigDealerID: line[2],
		ProgramCode:  line[5],
		Status:       "Terminated",
		Date:         line[9],
	}
	return currentVehicle
}

func createActiveVehicleReport(line []string) VehicleData {
	currentVehicle := VehicleData{
		VIN:          line[1],
		OrigDealerID: line[2],
		ProgramCode:  line[8],
		Status:       "Active",
		Date:         line[9],
	}
	return currentVehicle
}

// BuildReport creates a vehicle data report for Utilization.
func BuildReport(filename string) VehicleReport {
	lines, err := ReadCsv(filename)

	if err != nil {
		panic(err)
	}

	vr := VehicleReport{}

	for index, line := range lines {

		if index == 0 {
			// skip header line
			continue
		}

		if strings.Contains(filename, "active") {
			vehicle := createActiveVehicleReport(line)
			vr.addVehicle(vehicle)
		} else {
			vehicle := createTerminatedVehicleReport(line)
			vr.addVehicle(vehicle)
		}
	}

	resp, err := json.Marshal(vr)
	fmt.Println("VVV", string(resp))
	fmt.Println("VR", vr)
	return vr
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

	csvData, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return csvData, nil
}
