package core

// Contains the CPU registers used by the Game Boy.
type Registers struct {
	A  uint8
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	H  uint8
	L  uint8
	F  Flags
	PC uint16
	SP uint16
}

// Returns the value of the word register AF.
func (rs *Registers) AF() uint16 {
	return (uint16(rs.A) << 8) | uint16(rs.F.Encode())
}

// Sets the value of the word register AF.
func (rs *Registers) SetAF(v uint16) {
	rs.A = uint8(v >> 8)
	rs.F.Decode(uint8((v & 0xFF)))
}

// Returns the value of the word register BC.
func (rs *Registers) BC() uint16 {
	return (uint16(rs.B) << 8) | uint16(rs.C)
}

// Sets the value of the word register BC.
func (rs *Registers) SetBC(v uint16) {
	rs.B = uint8(v >> 8)
	rs.C = uint8((v & 0xFF))
}

// Returns the value of the word register DE.
func (rs *Registers) DE() uint16 {
	return (uint16(rs.D) << 8) | uint16(rs.E)
}

// Sets the value of the word register DE.
func (rs *Registers) SetDE(v uint16) {
	rs.D = uint8(v >> 8)
	rs.E = uint8((v & 0xFF))
}

// Returns the value of the word register HL.
func (rs *Registers) HL() uint16 {
	return (uint16(rs.H) << 8) | uint16(rs.L)
}

// Sets the value of the word register HL.
func (rs *Registers) SetHL(v uint16) {
	rs.H = uint8(v >> 8)
	rs.L = uint8((v & 0xFF))
}
