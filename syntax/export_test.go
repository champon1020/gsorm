package syntax

// Exported values which is declared in from.go.
var (
	FromName     = (*From).name
	FromAddTable = (*From).addTable
)

// Exported values which is declared in insert.go.
var (
	InsertQuery     = (*Insert).query
	InsertAddTable  = (*Insert).addTable
	InsertAddColumn = (*Insert).addColumn
)

// Exported values which is declared in limit.go.
var (
	LimitName = (*Limit).name
)

// Exported values which is declared in offset.go.
var (
	OffsetName = (*Offset).name
)

// Exported values which is declared in orderby.go.
var (
	OrderByName = (*OrderBy).name
)

// Exported values which is declared in select.go.
var (
	SelectQuery     = (*Select).query
	SelectAddColumn = (*Select).addColumn
)

// Exported values which is declared in set.go.
var (
	SetName  = (*Set).name
	SetAddEq = (*Set).addEq
)

// Exported values which is declared in update.go.
var (
	UpdateQuery     = (*Update).query
	UpdateAddTable  = (*Update).addTable
	UpdateAddColumn = (*Update).addColumns
)

// Exported values which is declared in values.go.
var (
	ValuesName      = (*Values).name
	ValuesAddColumn = (*Values).addColumn
)

// Exported values which is declared in where.go.
var (
	WhereName    = (*Where).name
	AndName      = (*And).name
	OrName       = (*Or).name
	BuildStmtSet = buildStmtSet
)
