package file

import (
	"encoding/csv"
	"fmt"
	"os"
)

type VehicleData struct {
	VIN          string
	OrigDealerID string
	ProgramCode  string
	Date         string
	Status       string
}

type VehicleReport struct {
	Vehicles []VehicleData
}

func (vehicle *VehicleReport) AddItem(car VehicleData) {
	vehicle.Vehicles = append(vehicle.Vehicles, car)
	fmt.Println("vee", vehicle.Vehicles)
}

func createTerminatedVehicleReport(line []string) []VehicleReport {

	currentVehicle := VehicleData{
		VIN:          line[1],
		OrigDealerID: line[2],
		ProgramCode:  line[5],
		Status:       "Terminated",
		Date:         line[9],
	}

	car := VehicleReport{}
	car.AddItem(currentVehicle)
	fmt.Println(car)
	return []VehicleReport{}
}

func BuildReport(filename string) {
	lines, err := ReadCsv(filename)

	if err != nil {
		panic(err)
	}

	for index, line := range lines {
		if index == 0 {
			// skip header line
			continue
		}

		// if strings.Contains(filename, "active") {
		// 	append(vehicleReport, createActiveVehicleReport(line))
		// } else {
		createTerminatedVehicleReport(line)
		// }
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

	csvData, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return csvData, nil
}
