//go:build baremetal && serial.none && !serial.mem
// +build baremetal,serial.none,!serial.mem

package machine

// Serial is a null device: writes to it are ignored.
var Serial = NullSerial{}

func InitSerial() {
	Serial.Configure(UARTConfig{})
}
