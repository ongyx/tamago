package tamago

type Clock struct {
	// t cycle = m cycle * 4
	m, t int
}

func (c *Clock) advance(steps int) {
	c.m += steps
	c.t += steps * 4
}

func (c *Clock) reset() {
	c.m = 0
	c.t = 0
}
