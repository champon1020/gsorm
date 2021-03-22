package internal

type Type string

var (
	Int    Type = "INT"
	Float  Type = "FLOAT"
	String Type = "VARCHAR(64)"
	Time   Type = "DATE"
	Bool   Type = "INT(1)"
)
