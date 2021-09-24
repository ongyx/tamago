package tamago

const (
	VBlank uint8 = 1 << iota
	LCDStat
	Timer
	Serial
	Joypad
)

var priority = [...]uint8{VBlank, LCDStat, Timer, Serial, Joypad}

// An interrupt is an event that is handled before the next instruction executes.
type Interrupt struct {
	master             bool
	enabled, requested uint8
}

func NewInterrupt() *Interrupt {
	return &Interrupt{master: true}
}

func (ir *Interrupt) todo() uint8 {
	flags := ir.enabled & ir.requested

	if ir.master && flags > 0 {
		for _, flag := range priority {
			if (flags & flag) != 0 {
				ir.requested &^= flag
				ir.master = false
			}
		}
		return flags
	} else {
		return 0
	}

}
