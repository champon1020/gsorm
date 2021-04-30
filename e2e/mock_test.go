package e2e_test

import (
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/statement"
	"github.com/google/go-cmp/cmp"
)

func TestSelectAllEmployees(t *testing.T) {
	testCases := []struct {
		Stmt   *statement.SelectStmt
		Result *[]Employee
	}{
		{
			mgorm.Select(db, "*").From("employees").(*statement.SelectStmt),
			&[]Employee{
				{
					EmpNo:     10001,
					BirthDate: time.Date(1953, time.September, 2, 0, 0, 0, 0, time.UTC),
					FirstName: "Georgi",
					LastName:  "Facello",
					Gender:    "M",
					HireDate:  time.Date(1986, time.June, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10002,
					BirthDate: time.Date(1964, time.June, 2, 0, 0, 0, 0, time.UTC),
					FirstName: "Bezalel",
					LastName:  "Simmel",
					Gender:    "F",
					HireDate:  time.Date(1985, time.November, 21, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10003,
					BirthDate: time.Date(1959, time.December, 3, 0, 0, 0, 0, time.UTC),
					FirstName: "Parto",
					LastName:  "Bamford",
					Gender:    "M",
					HireDate:  time.Date(1986, time.August, 28, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10004,
					BirthDate: time.Date(1954, time.May, 1, 0, 0, 0, 0, time.UTC),
					FirstName: "Chirstian",
					LastName:  "Koblick",
					Gender:    "M",
					HireDate:  time.Date(1986, time.December, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10005,
					BirthDate: time.Date(1955, time.January, 21, 0, 0, 0, 0, time.UTC),
					FirstName: "Kyoichi",
					LastName:  "Maliniak",
					Gender:    "M",
					HireDate:  time.Date(1989, time.September, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10006,
					BirthDate: time.Date(1953, time.April, 20, 0, 0, 0, 0, time.UTC),
					FirstName: "Anneke",
					LastName:  "Preusig",
					Gender:    "F",
					HireDate:  time.Date(1989, time.June, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10007,
					BirthDate: time.Date(1957, time.May, 23, 0, 0, 0, 0, time.UTC),
					FirstName: "Tzvetan",
					LastName:  "Zielinski",
					Gender:    "F",
					HireDate:  time.Date(1989, time.February, 10, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10008,
					BirthDate: time.Date(1958, time.February, 19, 0, 0, 0, 0, time.UTC),
					FirstName: "Saniya",
					LastName:  "Kalloufi",
					Gender:    "M",
					HireDate:  time.Date(1994, time.September, 15, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10009,
					BirthDate: time.Date(1952, time.April, 19, 0, 0, 0, 0, time.UTC),
					FirstName: "Sumant",
					LastName:  "Peac",
					Gender:    "F",
					HireDate:  time.Date(1985, time.February, 18, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:     10010,
					BirthDate: time.Date(1963, time.June, 1, 0, 0, 0, 0, time.UTC),
					FirstName: "Duangkaew",
					LastName:  "Piveteau",
					Gender:    "F",
					HireDate:  time.Date(1989, time.August, 24, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]Employee)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestSelectAllSalaries(t *testing.T) {
	testCases := []struct {
		Stmt   *statement.SelectStmt
		Result *[]Salary
	}{
		{
			mgorm.Select(db, "*").From("salaries").(*statement.SelectStmt),
			&[]Salary{
				{
					EmpNo:    10001,
					Salary:   60117,
					FromDate: time.Date(1986, time.June, 26, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1987, time.June, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10001,
					Salary:   62102,
					FromDate: time.Date(1987, time.June, 26, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1988, time.June, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10002,
					Salary:   65828,
					FromDate: time.Date(1996, time.August, 3, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1997, time.August, 3, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10002,
					Salary:   65909,
					FromDate: time.Date(1997, time.August, 3, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1998, time.August, 3, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10004,
					Salary:   40054,
					FromDate: time.Date(1986, time.December, 1, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1987, time.December, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10004,
					Salary:   42283,
					FromDate: time.Date(1987, time.December, 1, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1988, time.November, 30, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10007,
					Salary:   56724,
					FromDate: time.Date(1989, time.February, 10, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1990, time.February, 10, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10007,
					Salary:   60740,
					FromDate: time.Date(1990, time.February, 10, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1991, time.February, 10, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10009,
					Salary:   60929,
					FromDate: time.Date(1985, time.February, 18, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1986, time.February, 18, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10009,
					Salary:   64604,
					FromDate: time.Date(1986, time.February, 18, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1987, time.February, 18, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]Salary)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}

func TestSelectAllTitles(t *testing.T) {
	testCases := []struct {
		Stmt   *statement.SelectStmt
		Result *[]Title
	}{
		{
			mgorm.Select(db, "*").From("titles").(*statement.SelectStmt),
			&[]Title{
				{
					EmpNo:    10001,
					Title:    "Senior Engineer",
					FromDate: time.Date(1986, time.June, 26, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10002,
					Title:    "Staff",
					FromDate: time.Date(1996, time.August, 3, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10003,
					Title:    "Senior Engineer",
					FromDate: time.Date(1995, time.December, 3, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10004,
					Title:    "Engineer",
					FromDate: time.Date(1986, time.December, 1, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1995, time.December, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10004,
					Title:    "Senior Engineer",
					FromDate: time.Date(1995, time.December, 1, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10005,
					Title:    "Senior Staff",
					FromDate: time.Date(1996, time.September, 12, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10005,
					Title:    "Staff",
					FromDate: time.Date(1989, time.September, 12, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1996, time.September, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10006,
					Title:    "Senior Engineer",
					FromDate: time.Date(1990, time.August, 5, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10007,
					Title:    "Senior Staff",
					FromDate: time.Date(1996, time.February, 11, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(9999, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					EmpNo:    10007,
					Title:    "Staff",
					FromDate: time.Date(1989, time.February, 10, 0, 0, 0, 0, time.UTC),
					ToDate:   time.Date(1996, time.February, 11, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]Title)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			t.Errorf("Differs: (-want +got)\n%s", diff)
		}
	}
}
