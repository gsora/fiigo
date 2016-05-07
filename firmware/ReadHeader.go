package firmware

import (
	"bufio"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"time"
)

var fw *os.File
var err error
var fwreader *bufio.Reader

const stdDate string = "200601021504"

// ReadHeader takes in input the path of a Fiio firmware file and reads header information
// into a Structure struct.
//
// This function will return the reader initialized, with the seek pointer
// already pointing at the file list.
func (f *Firmware) ReadHeader(fwPath string) error {

	fw, err = os.Open(fwPath)

	if err != nil {
		return err
	}

	f.File = fw

	fwreader = bufio.NewReader(fw)

	// read the first 4 bytes into "title"
	title := []byte{}
	for i := 0; i < 4; i++ {
		a, _ := fwreader.ReadByte()
		title = append(title, a)
	}

	// first check if the reader points to a Fiio firmware,
	// if not, return error
	if string(title) != "SFHI" {
		return errors.New("Error - the file you provided is not a valid Fiio firmware")
	}

	f.Reader = fwreader
	s := Structure{}
	s.HeaderStartToken = string(title)

	// from this point, the reader already have offset = 4

	// number of sectors
	fsSectorsSlice := seekAndRead(f.Reader, 4, false)
	fsSectors, _ := strconv.ParseInt("0x"+hex.EncodeToString(fsSectorsSlice), 0, 64)
	s.FilesystemSectors = fsSectors

	// header size
	fsHeaderSizeSlice := seekAndRead(f.Reader, 4, true)
	fsHeaderSize, _ := strconv.ParseInt("0x"+hex.EncodeToString(fsHeaderSizeSlice), 0, 64)
	s.HeaderSize = fsHeaderSize

	// garbage data, move 4 forward
	emptySeek(f.Reader, 4)

	// date of creation
	dcSlice := seekAndRead(f.Reader, 12, false)
	t, _ := time.Parse(stdDate, string(dcSlice))
	s.CreationTime = t

	// number of files in the filesystem
	fnSlice := seekAndRead(f.Reader, 4, true)
	numberOfFiles, _ := strconv.ParseInt("0x"+hex.EncodeToString(fnSlice), 0, 64)
	s.NumberOfFiles = numberOfFiles

	// crc is always present
	s.CRC = CRCPresent

	emptySeek(f.Reader, 16)

	// machine id
	machineIDSlice := seekAndRead(f.Reader, 8, false)
	s.MachineString = string(machineIDSlice)

	s.FilesystemSectorSize = 512
	s.FilesystemSizeBytes = s.FilesystemSectorSize * s.FilesystemSectors

	emptySeek(f.Reader, 1992)

	f.Header = s

	return nil
}
