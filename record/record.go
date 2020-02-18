package record

// VehicleData is a single car record.
type VehicleData struct {
	VIN          string
	OrigDealerID string
	ProgramCode  string
	Date         string
	Status       string
}

// CreateTerminatedVehicleEntry creates a terminated vehicle entry.
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
