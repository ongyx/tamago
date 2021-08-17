package tamago

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type CPU struct {
	// 16-bit address bus (0000 - FFFF)
	mem   [0xFFFF]uint8
	gpu   *GPU
	reg   *Register
	intr  *Interrupt
	clock *Clock
}

func NewCPU() *CPU {
	return &CPU{
		gpu:   &GPU{},
		reg:   &Register{},
		intr:  &Interrupt{},
		clock: &Clock{},
	}
}

func (cpu *CPU) Read(address uint16) uint8 {

	/*
		switch address {

		case 0xff00:
			// TODO
			return 0

		case 0xff04:
			return uint8(rand.Intn(math.MaxUInt8))

		case 0xff0f:
			return cpu.intr.flags

		case 0xff40:
			return cpu.gpu.control

		case 0xff42:
			return cpu.gpu.scrollY

		case 0xff43:
			return cpu.gpu.scrollX

		case 0xff44:
			return cpu.gpu.scanline

		case 0xffff:
			return cpu.intr.enable

		default:

		}
	*/

	return cpu.mem[address]
}

func (cpu *CPU) Write(address uint16, value uint8) {
	cpu.mem[address] = value
}

func (cpu *CPU) WriteShort(address uint16, value uint16) {
	setShort(&cpu.mem[address], &cpu.mem[address+1], value)
}
