//go:build baremetal && serial.none && serial.mem
// +build baremetal,serial.none,serial.mem

package machine

import "unsafe"

//go:extern _serial_start
var _serial_start [0]byte

var (
	Serial   = memSerial{}
	printBuf = (*uintptr)(unsafe.Pointer(&_serial_start))
)

func InitSerial() {
}

func init() {
	*printBuf = uintptr(unsafe.Pointer(printBuf)) + 4
}

type memSerial struct {
}

func (memSerial) Buffered() int {
	return 0
}

func (memSerial) ReadByte() (byte, error) {
	return 0, errNoByte
}

func (ms memSerial) Write(p []byte) (n int, err error) {
	for _, b := range p {
		ms.WriteByte(b)
	}
	return len(p), nil
}

func (memSerial) WriteByte(b byte) error {
	*(*byte)(unsafe.Pointer(*printBuf)) = b
	*printBuf++
	return nil
}
