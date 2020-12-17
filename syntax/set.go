package syntax

// Set expression.
type Set struct {
	Eqs []Eq
}

func (s *Set) name() string {
	return "SET"
}

func (s *Set) addEq(lhs string, rhs interface{}) {
	e := NewEq(lhs, rhs)
	s.Eqs = append(s.Eqs, *e)
}

// Build make set statement set.
func (s *Set) Build() (*StmtSet, error) {
	ss := new(StmtSet)
	ss.WriteClause(s.name())
	for i, eq := range s.Eqs {
		if i != 0 {
			ss.WriteValue(",")
		}

		e, err := eq.Build()
		if err != nil {
			return nil, err
		}
		ss.WriteValue(e)
	}
	return ss, nil
}

// NewSet create new set object.
func NewSet(lhs []string, rhs []interface{}) (*Set, error) {
	if len(lhs) != len(rhs) {
		return nil, newError(ErrInvalid, "Length is different between lhs and rhs")
	}
	s := new(Set)
	for i := 0; i < len(lhs); i++ {
		s.addEq(lhs[i], rhs[i])
	}
	return s, nil
}
