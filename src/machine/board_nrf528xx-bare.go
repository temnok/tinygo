//go:build nrf528xx_bare
// +build nrf528xx_bare

package machine

const HasLowFrequencyCrystal = false

// UART pins
const (
	UART_TX_PIN Pin = NoPin
	UART_RX_PIN Pin = NoPin
)

// I2C pins (unused)
const (
	SDA_PIN = NoPin
	SCL_PIN = NoPin
)

// SPI pins (unused)
const (
	SPI0_SCK_PIN = NoPin
	SPI0_SDO_PIN = NoPin
	SPI0_SDI_PIN = NoPin
)
