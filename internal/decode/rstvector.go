package decode

const (
	RST0 RSTVector = 0x00
	RST1 RSTVector = 0x08
	RST2 RSTVector = 0x10
	RST3 RSTVector = 0x18
	RST4 RSTVector = 0x20
	RST5 RSTVector = 0x28
	RST6 RSTVector = 0x30
	RST7 RSTVector = 0x38
)

// An RST vector to jump to.
//
//go:generate stringer -type=RSTVector
type RSTVector uint8
