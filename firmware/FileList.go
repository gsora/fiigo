package firmware

import (
	"encoding/hex"
	"strconv"
)

// FileList return a FileEntry array
func (f *Firmware) FileList() {
	r := []FileEntry{}

	for i := 0; i < int(f.Header.NumberOfFiles); i++ {
		filePath := string(seekAndRead(f.Reader, 56, false))
		startSectorSlice := seekAndRead(f.Reader, 4, true)
		startSector, _ := strconv.ParseInt("0x"+hex.EncodeToString(startSectorSlice), 0, 64)
		fileSizeSlice := seekAndRead(f.Reader, 4, true)
		fileSize, _ := strconv.ParseInt("0x"+hex.EncodeToString(fileSizeSlice), 0, 64)
		r = append(r, FileEntry{Path: filePath, StartSector: startSector, Size: fileSize})
	}

	f.Files = r
}
