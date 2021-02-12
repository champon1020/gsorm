package syntax

import "strings"

// StmtSet is the statement set.
type StmtSet struct {
	Keyword string
	Value   string
	Parens  bool
}

// WriteKeyword write caluse to StmtSet.
func (ss *StmtSet) WriteKeyword(clause string) {
	if ss.Keyword != "" {
		ss.Keyword += " "
	}
	ss.Keyword += clause
}

// WriteValue write value to StmtSet.
func (ss *StmtSet) WriteValue(value string) {
	if ss.Value != "" && value != "," && value != ")" && !strings.HasSuffix(ss.Value, "(") {
		ss.Value += " "
	}
	ss.Value += value
}

// Build make sql string.
func (ss *StmtSet) Build() string {
	s := ss.Keyword
	if ss.Parens || ss.Value != "" {
		s += " "
	}
	if ss.Parens {
		s += "("
	}
	s += ss.Value
	if ss.Parens {
		s += ")"
	}
	return s
}
