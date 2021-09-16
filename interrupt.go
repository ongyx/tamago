package tamago

const (
	VBlank uint8 = 1 << iota
	LCDStat
	Timer
	Serial
	Joypad
)

type handler struct {
	flag uint8
	fn   func()
}

// An interrupt is an event that is handled before the next instruction executes.
type Interrupt struct {
	master             bool
	enabled, requested uint8
	handlers           []handler
}

func NewInterrupt() *Interrupt {
	return &Interrupt{}
}

func (ir *Interrupt) step(s *State) {
	flags := ir.enabled & ir.requested

	if ir.master && flags > 0 {
		for _, handler := range ir.handlers {
			if (flags & handler.flag) != 0 {
				ir.requested &^= handler.flag
				handler.fn()
			}
		}
	}
}
