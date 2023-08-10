package serial_comunication

import (
	"log"

	"github.com/tarm/serial"
)

func Teste() {
	c := &serial.Config{Name: "COM45", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	s.Write([]byte("200f200b"))
}

func service() *serial.Port {
	c := &serial.Config{Name: "COM45", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	return s

}

func SerialSender(recipe string) {
	service().Write([]byte(recipe))
}
