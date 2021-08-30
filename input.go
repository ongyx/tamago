package tamago

type button uint8

const (
	BtnA button = 1 << iota
	BtnB
	BtnSelect
	BtnStart
	BtnRight
	BtnLeft
	BtnUp
	BtnDown
)

type Input struct {
	// mode will be true for the action buttons,
	// false for the directional buttons.
	action  bool
	buttons button
}

func NewInput() *Input {
	// The bits are 1 if there are no keypresses, and 0 if there is a keypress.
	return &Input{action: true, buttons: button(0xFF)}
}

func (i *Input) Action(m bool) {
	i.action = m
}

func (i *Input) Press(btn button) {
	i.buttons &^= btn
}

func (i *Input) Release(btn button) {
	i.buttons |= btn
}

func (i *Input) Poll() uint8 {
	btns := uint8(i.buttons)

	if i.action {
		return (btns & 0x0F) | 0x10
	} else {
		return ((btns & 0xF0) >> 4) | 0x20
	}
}
