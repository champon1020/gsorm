package mig

import "github.com/champon1020/mgorm/internal"

type TColumn struct {
	Name          string
	Type          string
	NotNull       bool
	AutoIncrement bool
	Default       interface{}
}

func (c *TColumn) Build() (string, error) {
	var s internal.SQL

	s.Write(c.Name)
	s.Write(c.Type)
	if c.NotNull {
		s.Write("NOT NULL")
	}
	if c.AutoIncrement {
		s.Write("AUTO INCREMENT")
	}
	if c.Default != nil {
		str, err := internal.ToString(c.Default, true)
		if err != nil {
			return "", err
		}
		s.Write(str)
	}

	return string(s), nil
}
