package tamago

const (
	IRVBlank uint8 = 1 << iota
	IRLCDStat
	IRTimer
	IRSerial
	IRJoypad
)

// An interrupt is an event that is handled before the next instruction executes.
type Interrupt struct {
	master, enable bool
	flags          uint8
}

func NewInterrupt() *Interrupt {
	return &Interrupt{}
}

func (i *Interrupt) Enable() {
	i.master = true
}

func (i *Interrupt) Disable() {
	i.master = false
}

func (i *Interrupt) Enabled() bool {
	return i.master && i.enable
}

func (i *Interrupt) Handle(s *State) {
	// TODO
}
