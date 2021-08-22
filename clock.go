package tamago

// Clock keeps track of how many transistor cycles have run so far.
type Clock struct {
	t int
}

func NewClock() *Clock {
	return &Clock{}
}

// Step forward the clock by m machine cycles.
func (c *Clock) Step(m int) {
	c.t += m * 4
}

// Reset the clock to 0.
func (c *Clock) Reset() {
	c.t = 0
}
