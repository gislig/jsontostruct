package device

// Bios - Some comment
type Bios struct {
	BiosVersion string  `json:"bios_version"`
	BiosName    string  `json:"bios_name"`
	BiosGUID    string  `json:"bios_guid"`
	DeviceID    uint    `json:"device_id"`
	BiosPrice   int     `json:"bios_price"`
	BiosLong    int64   `json:"bios_long"`
	BiosFloat   float64 `json:"bios_float"`
	BiosBool    bool    `json:"bios_bool"`
}
