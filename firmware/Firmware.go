package firmware

import (
	"bufio"
	"os"
)

// Firmware represent the firmware file
type Firmware struct {
	File   *os.File
	Reader *bufio.Reader
	Path   string
	Header Structure
}
