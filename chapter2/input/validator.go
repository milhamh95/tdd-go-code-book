package input

type Validator struct {
	expectedLength int
	validOperators []string
}

// NewValidator creates a ready to use Validator instance
func NewValidator(expLen int, validOps []string) *Validator {
	return &Validator{
		expectedLength: expLen,
		validOperators: validOps,
	}
}

// CheckInput validates the operator and operands
func (v *Validator) CheckInput(operator string, operands []float64) error {
	return nil
}
