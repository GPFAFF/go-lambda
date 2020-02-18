package record

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

// CreateTerminatedVehicleEntry creates an active vehicle entry.
func CreateTerminatedVehicleEntry(line []string) VehicleData {
	return VehicleData{
		VIN:          line[1],
		OrigDealerID: line[2],
		ProgramCode:  line[5],
		Status:       "Terminated",
		Date:         line[9],
	}
}

// CreateActiveVehicleEntry creates an active vehicle entry.
func CreateActiveVehicleEntry(line []string) VehicleData {
	return VehicleData{
		VIN:          line[1],
		OrigDealerID: line[2],
		ProgramCode:  line[8],
		Status:       "Active",
		Date:         line[9],
	}
}
