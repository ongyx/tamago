package tamago

const (
	BtnRight uint8 = 1 << iota
	BtnLeft
	BtnUp
	BtnDown
	BtnA
	BtnB
	BtnSelect
	BtnStart

	SelectDir uint8 = 0x10
	SelectAct uint8 = 0x20
)

type Input struct {
	btns, sel uint8
}

func NewInput() *Input {
	// The bits are 1 if there are no keypresses, and 0 if there is a keypress.
	// Both the directional and action buttons are selected.
	return &Input{btns: uint8(0xFF)}
}

func (i *Input) Press(btn uint8) {
	i.btns &^= btn
}

func (i *Input) Release(btn uint8) {
	i.btns |= btn
}

func (i *Input) Select(v uint8) {
	i.sel = v & (SelectDir | SelectAct)
}

func (i *Input) Poll() uint8 {

	var v uint8

	if (i.sel & SelectDir) > 0 {
		v |= (i.btns & 0xf)
	}

	if (i.sel & SelectAct) > 0 {
		v |= (i.btns >> 4)
	}

	return v | i.sel
}
