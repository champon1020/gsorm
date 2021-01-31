package dockertest_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/internal"
	"github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/ory/dockertest/v3"
)

const TIMEOUT = 20 * time.Second

var db *mgorm.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	resource, err := pool.BuildAndRun("mysql-mock", "./image/Dockerfile", []string{})
	if err != nil {
		log.Fatalf("Could not start resource: %v", err)
	}

	if err := pool.Retry(func() error {
		db, err = mgorm.New(
			"mysql",
			fmt.Sprintf("root:toor@tcp(localhost:%s)/mock", resource.GetPort("3306/tcp")),
		)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	start := time.Now()
	// Ignore any errors which is printed by go-sql-driver/mysql.
	mysql.SetLogger(log.New(ioutil.Discard, "", 0))
	for {
		if db.Ping() == nil {
			break
		}
		now := time.Now()
		if (now.Sub(start)).Seconds() > float64(TIMEOUT) {
			log.Fatalf("Timeout")
		}
	}
	// Reset logger of go-sql-driver/mysql.
	mysql.SetLogger(log.New(os.Stderr, "[mysql] ", log.Ldate|log.Ltime|log.Lshortfile))

	m.Run()

	defer func() {
		db.Close()
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %v", err)
		}
	}()
}

type Model struct {
	ID           int
	Height       float32
	Gender       string
	DayOfTheWeek string
	Date1        time.Time `layout:"2006-01-02"`
	Date2        time.Time `layout:"2006-01-02 15:04:05"`
	FirstName    string
	LastName     string
}

func TestSelectAll(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *[]Model
	}{
		{
			mgorm.Select(db, "*").From("mock").(*mgorm.Stmt),
			&[]Model{
				{
					ID:           0,
					Height:       170.2,
					Gender:       "M",
					DayOfTheWeek: "Sun",
					Date1:        time.Date(2005, time.June, 15, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2006, time.September, 28, 20, 22, 22, 0, time.UTC),
					FirstName:    "Mick",
					LastName:     "Boult",
				},
				{
					ID:           1,
					Height:       160.3,
					Gender:       "F",
					DayOfTheWeek: "Mon",
					Date1:        time.Date(2016, time.August, 31, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(1984, time.December, 1, 16, 45, 45, 0, time.UTC),
					FirstName:    "Kathryn",
					LastName:     "Langshaw",
				},
				{
					ID:           2,
					Height:       180.4,
					Gender:       "M",
					DayOfTheWeek: "Tue",
					Date1:        time.Date(1997, time.August, 29, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2009, time.February, 22, 21, 29, 29, 0, time.UTC),
					FirstName:    "Alvin",
					LastName:     "O'Casey",
				},
				{
					ID:           3,
					Height:       179.8,
					Gender:       "M",
					DayOfTheWeek: "Wed",
					Date1:        time.Date(2009, time.February, 22, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(1987, time.March, 22, 5, 28, 28, 0, time.UTC),
					FirstName:    "Jake",
					LastName:     "Lonergan",
				},
				{
					ID:           4,
					Height:       155.2,
					Gender:       "F",
					DayOfTheWeek: "Thu",
					Date1:        time.Date(1999, time.October, 29, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(1999, time.December, 14, 19, 49, 49, 0, time.UTC),
					FirstName:    "Robin",
					LastName:     "Tate",
				},
				{
					ID:           5,
					Height:       169.1,
					Gender:       "M",
					DayOfTheWeek: "Fri",
					Date1:        time.Date(1987, time.June, 18, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2018, time.February, 2, 8, 58, 58, 0, time.UTC),
					FirstName:    "Clark",
					LastName:     "Rangeley",
				},
				{
					ID:           6,
					Height:       185.2,
					Gender:       "M",
					DayOfTheWeek: "Sat",
					Date1:        time.Date(2006, time.October, 9, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2010, time.September, 2, 15, 6, 6, 0, time.UTC),
					FirstName:    "Dom",
					LastName:     "Meehan",
				},
				{
					ID:           7,
					Height:       169.9,
					Gender:       "F",
					DayOfTheWeek: "Sun",
					Date1:        time.Date(2016, time.May, 2, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2007, time.October, 1, 11, 33, 33, 0, time.UTC),
					FirstName:    "Mavis",
					LastName:     "Pierce",
				},
				{
					ID:           8,
					Height:       168.8,
					Gender:       "F",
					DayOfTheWeek: "Mon",
					Date1:        time.Date(1988, time.December, 28, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2016, time.January, 12, 23, 39, 39, 0, time.UTC),
					FirstName:    "Bonnie",
					LastName:     "Reed",
				},
				{
					ID:           9,
					Height:       165.4,
					Gender:       "F",
					DayOfTheWeek: "Tue",
					Date1:        time.Date(1981, time.April, 17, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2016, time.October, 1, 0, 20, 20, 0, time.UTC),
					FirstName:    "Virginia",
					LastName:     "Riddell",
				},
				{
					ID:           10,
					Height:       170.3,
					Gender:       "M",
					DayOfTheWeek: "Wed",
					Date1:        time.Date(2013, time.October, 1, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(1989, time.December, 30, 22, 28, 28, 0, time.UTC),
					FirstName:    "Levi",
					LastName:     "Downes",
				},
				{
					ID:           11,
					Height:       150.5,
					Gender:       "F",
					DayOfTheWeek: "Thu",
					Date1:        time.Date(2013, time.August, 10, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(1994, time.April, 6, 2, 35, 35, 0, time.UTC),
					FirstName:    "Lucille",
					LastName:     "Coulthard",
				},
				{
					ID:           12,
					Height:       153.3,
					Gender:       "F",
					DayOfTheWeek: "Fri",
					Date1:        time.Date(2011, time.December, 29, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2000, time.March, 16, 4, 2, 2, 0, time.UTC),
					FirstName:    "Evangeline",
					LastName:     "Ballard",
				},
				{
					ID:           13,
					Height:       180.4,
					Gender:       "M",
					DayOfTheWeek: "Sat",
					Date1:        time.Date(2015, time.November, 14, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(1997, time.February, 19, 5, 39, 39, 0, time.UTC),
					FirstName:    "Deryck",
					LastName:     "Watson",
				},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]Model)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestWhere(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *[]Model
	}{
		{
			mgorm.Select(db, "*").
				From("mock").
				Where("day_of_the_week = ?", "Sun").(*mgorm.Stmt),
			&[]Model{
				{
					ID:           0,
					Height:       170.2,
					Gender:       "M",
					DayOfTheWeek: "Sun",
					Date1:        time.Date(2005, time.June, 15, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2006, time.September, 28, 20, 22, 22, 0, time.UTC),
					FirstName:    "Mick",
					LastName:     "Boult",
				},
				{
					ID:           7,
					Height:       169.9,
					Gender:       "F",
					DayOfTheWeek: "Sun",
					Date1:        time.Date(2016, time.May, 2, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2007, time.October, 1, 11, 33, 33, 0, time.UTC),
					FirstName:    "Mavis",
					LastName:     "Pierce",
				},
			},
		},
		{
			mgorm.Select(db, "height", "first_name").
				From("mock").
				Where("id > ?", 11).(*mgorm.Stmt),
			&[]Model{
				{
					Height:    153.3,
					FirstName: "Evangeline",
				},
				{
					Height:    180.4,
					FirstName: "Deryck",
				},
			},
		},
		{
			mgorm.Select(db, "height", "first_name").
				From("mock").
				Where("id > ? AND gender = ?", 11, "M").(*mgorm.Stmt),
			&[]Model{
				{
					Height:    180.4,
					FirstName: "Deryck",
				},
			},
		},
		{
			mgorm.Select(db, "id").
				From("mock").
				Where("id < ?", 8).
				And("first_name = ? OR last_name = ?", "Alvin", "Tate").(*mgorm.Stmt),
			&[]Model{
				{ID: 2}, {ID: 4},
			},
		},
		{
			mgorm.Select(db, "id").
				From("mock").
				Where("id > ?", 8).
				Or("first_name = ? AND last_name = ?", "Robin", "Tate").(*mgorm.Stmt),
			&[]Model{{ID: 4}, {ID: 9}, {ID: 10}, {ID: 11}, {ID: 12}, {ID: 13}},
		},
	}

	for i, testCase := range testCases {
		model := new([]Model)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestLimitOffsetOrderBy(t *testing.T) {
	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *[]Model
	}{
		{
			mgorm.Select(db, "id").
				From("mock").
				OrderBy("first_name", false).
				Limit(2).(*mgorm.Stmt),
			&[]Model{
				{ID: 2}, {ID: 8},
			},
		},
		{
			mgorm.Select(db, "id").
				From("mock").
				Limit(2).
				Offset(3).(*mgorm.Stmt),
			&[]Model{
				{ID: 3}, {ID: 4},
			},
		},
	}

	for i, testCase := range testCases {
		model := new([]Model)
		if err := testCase.Stmt.Query(model); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, model); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			internal.PrintTestDiff(t, diff)
		}
	}
}

func TestJoin(t *testing.T) {
	type Model2 struct {
		ID           int
		Height       float32
		Gender       string
		DayOfTheWeek string
		Date1        time.Time `layout:"2006-01-02"`
		Date2        time.Time `layout:"2006-01-02 15:04:05"`
		FirstName    string
		LastName     string
		ID2          int
		FromDate     time.Time `layout:"2006-01-02"`
		Company      string
		Country      string
	}

	testCases := []struct {
		Stmt   *mgorm.Stmt
		Result *[]Model2
	}{
		{
			mgorm.Select(db,
				"mock.id AS id",
				"height",
				"gender",
				"day_of_the_week",
				"date1",
				"date2",
				"first_name",
				"last_name",
				"mock2.id AS id2",
				"from_date",
				"company",
				"country").
				From("mock").
				Join("mock2").
				On("mock.id = mock2.id").
				OrderBy("id", false).
				Limit(2).(*mgorm.Stmt),
			&[]Model2{
				{
					ID:           1,
					Height:       160.3,
					Gender:       "F",
					DayOfTheWeek: "Mon",
					Date1:        time.Date(2016, time.August, 31, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(1984, time.December, 1, 16, 45, 45, 0, time.UTC),
					FirstName:    "Kathryn",
					LastName:     "Langshaw",
					ID2:          1,
					FromDate:     time.Date(2016, time.August, 31, 0, 0, 0, 0, time.UTC),
					Company:      "ABC Company",
					Country:      "Japan",
				},
				{
					ID:           2,
					Height:       180.4,
					Gender:       "M",
					DayOfTheWeek: "Tue",
					Date1:        time.Date(1997, time.August, 29, 0, 0, 0, 0, time.UTC),
					Date2:        time.Date(2009, time.February, 22, 21, 29, 29, 0, time.UTC),
					FirstName:    "Alvin",
					LastName:     "O'Casey",
					ID2:          2,
					FromDate:     time.Date(2009, time.February, 22, 0, 0, 0, 0, time.UTC),
					Company:      "DEF Holdings",
					Country:      "America",
				},
			},
		},
		{
			mgorm.Select(db,
				"mock.id AS id",
				"mock.first_name",
				"mock2.id AS id2",
				"company").
				From("mock").
				LeftJoin("mock2").
				On("mock.id = mock2.id").
				OrderBy("id", false).
				Limit(2).(*mgorm.Stmt),
			&[]Model2{
				{
					ID:        0,
					FirstName: "Mick",
				},
				{
					ID:        1,
					FirstName: "Kathryn",
					ID2:       1,
					Company:   "ABC Company",
				},
			},
		},
		{
			mgorm.Select(db,
				"mock.id AS id",
				"mock.first_name",
				"mock2.id AS id2",
				"mock2.company").
				From("mock").
				RightJoin("mock2").
				On("mock.id = mock2.id").
				OrderBy("id2", true).
				Limit(2).(*mgorm.Stmt),
			&[]Model2{
				{
					ID2:     14,
					Company: "JKL Film",
				},
				{
					ID:        3,
					FirstName: "Jake",
					ID2:       3,
					Company:   "GHI Corporation",
				},
			},
		},
	}

	for i, testCase := range testCases {
		model2 := new([]Model2)
		if err := testCase.Stmt.Query(model2); err != nil {
			t.Errorf("Error was occurred: %v", err)
			continue
		}
		if diff := cmp.Diff(testCase.Result, model2); diff != "" {
			t.Errorf("Got difference with sample %d", i)
			internal.PrintTestDiff(t, diff)
		}
	}
}
