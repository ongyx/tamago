package core

// Contains the general purpose (A, B, C, D, E, H, L, F, PC, SP) registers and special purpose (IE, IF) registers used by the CPU.
type Registers struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8
	H uint8
	L uint8
	F ALUFlags

	// Are interrupts master enabled?
	IME bool
	// The specific interrupts enabled.
	IE InterruptFlags
	// The specific interrupts requested.
	IR InterruptFlags

	PC uint16
	SP uint16
}

// Returns the value of the word register AF.
func (rs *Registers) AF() uint16 {
	return CombineWord(rs.A, rs.F.Encode())
}

// Sets the value of the word register AF.
func (rs *Registers) SetAF(v uint16) {
	hi, lo := SplitWord(v)
	rs.A = hi
	rs.F.Decode(lo)
}

// Returns the value of the word register BC.
func (rs *Registers) BC() uint16 {
	return CombineWord(rs.B, rs.C)
}

// Sets the value of the word register BC.
func (rs *Registers) SetBC(v uint16) {
	rs.B, rs.C = SplitWord(v)
}

// Returns the value of the word register DE.
func (rs *Registers) DE() uint16 {
	return CombineWord(rs.D, rs.E)
}

// Sets the value of the word register DE.
func (rs *Registers) SetDE(v uint16) {
	rs.D, rs.E = SplitWord(v)
}

// Returns the value of the word register HL.
func (rs *Registers) HL() uint16 {
	return CombineWord(rs.H, rs.L)
}

// Sets the value of the word register HL.
func (rs *Registers) SetHL(v uint16) {
	rs.H, rs.L = SplitWord(v)
}

// Checks if there is an enabled and requested interrupt, and returns the corresponding interrupt vector.
//
// If no interrupts are enabled and requested, [InterruptVectorNone] is returned.
func (rs *Registers) CheckInterrupt() InterruptVector {
	if rs.IE.VBlank && rs.IR.VBlank {
		rs.IR.VBlank = false
		return InterruptVectorVBlank
	} else if rs.IE.Stat && rs.IR.Stat {
		rs.IR.Stat = false
		return InterruptVectorStat
	} else if rs.IE.Timer && rs.IR.Timer {
		rs.IR.Timer = false
		return InterruptVectorTimer
	} else if rs.IE.Serial && rs.IR.Serial {
		rs.IR.Serial = false
		return InterruptVectorSerial
	} else if rs.IE.Joypad && rs.IR.Joypad {
		rs.IR.Joypad = false
		return InterruptVectorJoypad
	}

	return InterruptVectorNone
}
