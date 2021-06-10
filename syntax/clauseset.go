package syntax

import "strings"

// ClauseSet is the pair of clause keyword and its values.
// If Parens is true, ClauseSet would be enclosed by parentheses.
type ClauseSet struct {
	Keyword string
	Value   string
	Parens  bool
}

// WriteKeyword writes caluse keyword to ClauseSet.
func (ss *ClauseSet) WriteKeyword(clause string) {
	if ss.Keyword != "" {
		ss.Keyword += " "
	}
	ss.Keyword += clause
}

// WriteValue writes value to ClauseSet.
func (ss *ClauseSet) WriteValue(value string) {
	if ss.Value != "" && value != "," && value != ")" && !strings.HasSuffix(ss.Value, "(") {
		ss.Value += " "
	}
	ss.Value += value
}

// Build makes clause with string.
func (ss *ClauseSet) Build() string {
	s := ss.Keyword
	if s != "" && (ss.Parens || ss.Value != "") {
		s += " "
	}
	s += ss.BuildValue()
	return s
}

// BuildValue makes clause value with string.
func (ss *ClauseSet) BuildValue() string {
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
