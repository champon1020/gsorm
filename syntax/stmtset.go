package syntax

import "strings"

// StmtSet is the pair of clause keyword and its values.
// If Parens is true, StmtSet would be enclosed by parentheses.
type StmtSet struct {
	Keyword string
	Value   string
	Parens  bool
}

// WriteKeyword writes caluse keyword to StmtSet.
func (ss *StmtSet) WriteKeyword(clause string) {
	if ss.Keyword != "" {
		ss.Keyword += " "
	}
	ss.Keyword += clause
}

// WriteValue writes value to StmtSet.
func (ss *StmtSet) WriteValue(value string) {
	if ss.Value != "" && value != "," && value != ")" && !strings.HasSuffix(ss.Value, "(") {
		ss.Value += " "
	}
	ss.Value += value
}

// Build makes clause with string.
func (ss *StmtSet) Build() string {
	s := ss.Keyword
	if s != "" && (ss.Parens || ss.Value != "") {
		s += " "
	}
	s += ss.BuildValue()
	return s
}

// BuildValue makes clause value with string.
func (ss *StmtSet) BuildValue() string {
	var s string
	if ss.Parens {
		s += "("
	}
	s += ss.Value
	if ss.Parens {
		s += ")"
	}
	return s
}
