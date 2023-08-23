package value_object

type Grade string

const (
	E  Grade = "E"
	D  Grade = "D"
	C  Grade = "C"
	B  Grade = "B"
	A_ Grade = "A-"
	A  Grade = "A"
)

type Score struct {
	Value float64
}

func NewScore(value float64) (*Score, error) {
	return &Score{Value: value}, nil
}

func (s *Score) CalculateGrade() Grade {
	switch {
	case s.Value < 1:
		return E
	case s.Value >= 1 && s.Value <= 1.5:
		return D
	case s.Value > 1.5 && s.Value <= 2.5:
		return C
	case s.Value > 2.5 && s.Value < 3:
		return B
	case s.Value >= 3 && s.Value <= 3.25:
		return A_
	default:
		return A
	}
}
