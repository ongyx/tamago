package tamago

// An interrupt is an event.
type Interrupt struct {
	master, enable, flags uint8
}
