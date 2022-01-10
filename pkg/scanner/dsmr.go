package scanner

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"go.bug.st/serial"
)

type state int
type TelegramCallback func([]byte, uint16)

const (
	idle state = iota
	telegram
	checksum
)

type DSMRScanner struct {
	port             serial.Port
	state            state
	unprocessed      []byte
	currentTelegram  []byte
	currentChecksum  []byte
	telegramCallback TelegramCallback
}

func NewDSMRScanner(port serial.Port, callback TelegramCallback) *DSMRScanner {
	s := &DSMRScanner{
		port:             port,
		state:            idle,
		unprocessed:      []byte{},
		currentTelegram:  []byte{},
		currentChecksum:  []byte{},
		telegramCallback: callback,
	}

	return s
}

func (s *DSMRScanner) readSerial() error {
	for {
		data := make([]byte, 4096)
		n, err := s.port.Read(data)
		if err != nil {
			return fmt.Errorf("error reading from serial port: %w", err)
		}
		if n > 0 {
			s.unprocessed = append(s.unprocessed, data[:n]...)
		} else {
			return nil
		}
	}
}

func (s *DSMRScanner) tokenize() {
	for len(s.unprocessed) > 0 {
		switch s.state {
		case idle:
			switch {
			// starts with '/'?
			case bytes.HasPrefix(s.unprocessed, []byte{47}):
				s.currentTelegram = s.unprocessed[:1]
				s.unprocessed = s.unprocessed[1:]
				s.state = telegram
			// '\n/' (ZIV meters)
			case bytes.HasPrefix(s.unprocessed, []byte{10, 47}):
				s.currentTelegram = s.unprocessed[:2]
				s.unprocessed = s.unprocessed[2:]
				s.state = telegram
			// HELP! unknown input. Discard a byte and stay idle
			default:
				s.unprocessed = s.unprocessed[1:]
			}
		case telegram:
			idx := bytes.IndexByte(s.unprocessed, 33) // !
			if idx == -1 {
				s.currentTelegram = append(s.currentTelegram, s.unprocessed...)
				s.unprocessed = []byte{}
			} else {
				s.state = checksum
				s.currentTelegram = append(s.currentTelegram, s.unprocessed[:idx]...)
				if len(s.unprocessed) > idx {
					s.unprocessed = s.unprocessed[idx+1:]
				} else {
					s.unprocessed = []byte{}
				}
			}
		case checksum:
			idx := bytes.Index(s.unprocessed, []byte{13, 10}) // CR + LF
			if idx == -1 {
				s.currentChecksum = append(s.currentChecksum, s.unprocessed...)
				s.unprocessed = []byte{}
			} else {
				s.currentChecksum = append(s.currentChecksum, s.unprocessed[:idx]...)
				s.telegramCallback(s.currentTelegram, binary.BigEndian.Uint16(s.currentChecksum))
				s.currentTelegram = []byte{}
				s.currentChecksum = []byte{}
				if len(s.unprocessed) > idx+1 {
					s.unprocessed = s.unprocessed[idx+2:]
				} else {
					s.unprocessed = []byte{}
				}
				s.state = idle
			}
		}
	}
}

func (s *DSMRScanner) ReadLoop() {
	for {
		s.readSerial()
		s.tokenize()
	}
}
