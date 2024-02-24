package errors

type ConstantError string

func (e ConstantError) Error() string { return string(e) }

// constant errors because i don't like sentinel errors
const (
	ErrNoOrder = ConstantError("Unit has no order")
)
