package core

const (
	InterruptVectorNone InterruptVector = 0x00

	InterruptVectorVBlank InterruptVector = 0x40 + (iota * 0x8)
	InterruptVectorLCD
	InterruptVectorTimer
	InterruptVectorSerial
	InterruptVectorJoypad
)

// An address to jump to after the corresponding interrupt is triggered.
//
//go:generate stringer -type=InterruptVector
type InterruptVector uint16
