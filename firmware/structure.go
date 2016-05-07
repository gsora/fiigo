package firmware

import "time"

// CRCMode represent whether or not CRC is present at the end of the firmware
//go:generate stringer -type=CRCMode
type CRCMode int64

const (
	// NoCRC = CRC is not present
	NoCRC CRCMode = iota
	// CRCPresent = CRC is present
	CRCPresent
)

// Structure represents the firmware header
type Structure struct {
	HeaderStartToken     string    // offset 0x0, length 4 - confirmed
	FilesystemSectors    int64     // offset 0x4, length 4 - unsure
	HeaderSize           int64     // offset 0x8, length 4 - unsure
	CreationTime         time.Time // offset 0x10, length 12 - confirmed
	NumberOfFiles        int64     // offset 0x1c, length 4 - confirmed
	CRC                  CRCMode   // offset 0x20, length 4 - confirmed but always true
	MachineString        string    // offset 0x30, length 8 - confirmed
	FilesystemSizeBytes  int64     // not sure
	FilesystemSectorSize int64     // confirmed, always 512
}
