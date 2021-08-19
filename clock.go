package tamago

// Clock keeps track of how many machine cycles have run so far.
type Clock struct {
	m, t int
}

// Step forward the clock by n machine cycles.
func (c *Clock) Step(n int) {
	c.m += n
	c.t += n * 4
}

// Reset the clock to 0.
func (c *Clock) Reset() {
	c.m = 0
	c.t = 0
}
