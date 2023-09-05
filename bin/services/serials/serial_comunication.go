package serials

import (
	"log"

	"github.com/tarm/serial"
)

var SerialConnect *serial.Port

func Connection() {
	c := &serial.Config{Name: "COM45", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	SerialConnect = s

}
