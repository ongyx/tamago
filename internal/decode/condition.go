package decode

const (
	ConditionNone Condition = iota
	ConditionZero
	ConditionNotZero
	ConditionCarry
	ConditionNotCarry
)

// Represents a condition for a JP, JR, CALL, or RET instruction to take effect.
//
//go:generate stringer -type=Condition
type Condition uint8
