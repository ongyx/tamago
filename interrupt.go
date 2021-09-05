package tamago

type iflag uint8

const (
	VBlank iflag = 1 << iota
	LCDStat
	Timer
	Serial
	Joypad
)

// A function that is called when the registered interrupt fires.
type Handler func(iflag uint8)

// An interrupt is an event that is handled before the next instruction executes.
type Interrupt struct {
	master, enable uint8
	flags          iflag
	handler        map[iflag]Handler
}

func NewInterrupt() *Interrupt {
	return &Interrupt{handler: make(map[iflag]Handler)}
}

// Register a handler for a callback.
func (ir *Interrupt) Register(f iflag, h Handler) {
	ir.handler[f] = h
}
