package main

import (
	"fmt"
	"time"

	"github.com/jonkerj/go-iec62056/pkg/parser"
	"github.com/jonkerj/go-iec62056/pkg/scanner"
	"go.bug.st/serial"
)

func telegramCallback(datagram []byte, checksum uint16) {
	telegram, err := parser.Parse(datagram)
	if err != nil {
		panic(fmt.Sprintf("parse error: %v", err))
	}

	fmt.Printf("telegram: %v\n", telegram)
	for _, obj := range telegram.Objects {
		fmt.Printf("obj: %v\n", obj)
	}
}

func main() {
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		panic(err)
	}
	port.SetReadTimeout(10 * time.Millisecond)

	scanner := scanner.NewDSMRScanner(port, telegramCallback)
	scanner.ReadLoop()
}
