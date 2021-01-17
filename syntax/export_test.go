package syntax

// Exported values which is declared int from.go.
var (
	FromName     = (*From).name
	FromAddTable = (*From).addTable
)

// Exported values which is declared int insert.go.
var (
	InsertQuery     = (*Insert).query
	InsertAddTable  = (*Insert).addTable
	InsertAddColumn = (*Insert).addColumn
)

// Exported values which is declared int select.go.
var (
	SelectQuery     = (*Select).query
	SelectAddColumn = (*Select).addColumn
)

// Exported values which is declared int set.go.
var (
	SetName  = (*Set).name
	SetAddEq = (*Set).addEq
)

// Exported values which is declared int update.go.
var (
	UpdateQuery     = (*Update).query
	UpdateAddTable  = (*Update).addTable
	UpdateAddColumn = (*Update).addColumns
)

// Exported values which is declared int values.go.
var (
	ValuesName      = (*Values).name
	ValuesAddColumn = (*Values).addColumn
)

// Exported values which is declared int where.go.
var (
	WhereName    = (*Where).name
	AndName      = (*And).name
	OrName       = (*Or).name
	BuildStmtSet = buildStmtSet
)