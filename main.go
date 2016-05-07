package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gsora/fiigo/firmware"
)

var fwPath string

func main() {
	flag.StringVar(&fwPath, "fw", "", "path of the firmware file")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	f := new(firmware.Firmware)
	f.ReadHeader(fwPath)
	fmt.Println(f.Header)
	f.FileList()
	f.Unpack("./fw")

}
