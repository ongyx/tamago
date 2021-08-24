package tamago

type iflag uint8

const (
	VBlank iflag = 1 << iota
	LCDStat
	Timer
	Serial
	Joypad
)

// An interrupt is an event that is handled before the next instruction executes.
type Interrupt struct {
	master, enable bool
	flags          iflag
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

func (i *Interrupt) Enabled() {
	return i.master && i.enable
}

func (i *Interrupt) Handle(s *State) {
	// TODO
}
