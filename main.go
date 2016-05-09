package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/fatih/structs"
	"github.com/gsora/fiigo/firmware"
)

var fwPath string
var unpack string
var pack string
var info bool

func main() {
	flag.StringVar(&fwPath, "firmware", "", "path of the firmware file")
	flag.StringVar(&unpack, "unpack", "", "unpack the firmware in the given directory")
	flag.StringVar(&pack, "pack", "", "pack the firmware in the given directory")
	flag.BoolVar(&info, "info", false, "print header of the given firmware")
	flag.Parse()

	if (unpack == "" && pack == "" && info == false) || fwPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	f := new(firmware.Firmware)
	f.ReadHeader(fwPath)
	f.FileList()

	if unpack != "" {
		f.Unpack(unpack)
	}

	if pack != "" {
		fmt.Println("Pack is not implemented yet!")
	}

	if info != false {
		title := "Firmware header for file " + f.Path
		fmt.Println(title)
		for i := 0; i < len(title); i++ {
			fmt.Printf("-")
		}
		fmt.Println("")
		for key, value := range structs.Map(f.Header) {
			if key == "CRC" {
				if value == firmware.CRCPresent {
					color.Green(key + " = present")
				} else if value == firmware.NoCRC {
					color.Green(key + " = not present")
				}
			} else {
				switch v := value.(type) {
				case string:
					color.Green("%s = %s", key, v)
				case int64:
					color.Green("%s = %d", key, v)
				}

			}
		}
	}

}
