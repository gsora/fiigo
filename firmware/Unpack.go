package firmware

import (
	"log"
	"os"
	"strings"
)

// Unpack unpack the firmware in the destination directory
func (f *Firmware) Unpack(dest string) {
	if _, err := os.Stat("./conf/app.ini"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dest, 0755)
		}
	}

	for _, i := range f.Files {
		createTree(dest + "/" + strings.Replace(i.Path, "\\", "/", -1))
	}
}

func createTree(s string) {
	err := os.MkdirAll(getPath(s), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getPath(s string) string {
	st := strings.Split(s, "/")

	if len(st) == 1 {
		return st[0]
	}

	r := ""
	for i := 0; i < len(st)-1; i++ {
		if i != 0 {
			r = r + "/" + st[i]
		} else {
			r = r + st[i]
		}
	}
	return r
}
