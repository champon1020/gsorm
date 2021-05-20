package e2e_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/champon1020/mgorm"
	"github.com/champon1020/mgorm/interfaces/domain"

	"github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
)

const TIMEOUT = 20 * time.Second

var db domain.DB

func TestMain(m *testing.M) {
	if os.Getenv("MGORM_CI") == "true" {
		RunTest(m)
		return
	}
	RunTestWithBuild(m)
}

func RunTest(m *testing.M) {
	var err error
	db, err = mgorm.Open("mysql", "root:toor@tcp(localhost:3306)/employees?parseTime=true")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	defer func() {
		db.Close()
	}()

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
}

func RunTestWithBuild(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	// Build container.
	resource, err := pool.BuildAndRun("mysql-mock", "./image/Dockerfile", []string{})
	if err != nil {
		log.Fatalf("Could not start resource: %v", err)
	}

	// Connect to database.
	if err = pool.Retry(func() error {
		db, err = mgorm.Open(
			"mysql",
			fmt.Sprintf("root:toor@tcp(localhost:%s)/employees?parseTime=true", resource.GetPort("3306/tcp")),
		)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	defer func() {
		db.Close()
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %v", err)
		}
	}()

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
}

type Employee struct {
	EmpNo     int
	BirthDate time.Time `mgorm:"layout=2006-01-02"`
	FirstName string
	LastName  string
	Gender    string
	HireDate  time.Time `mgorm:"layout=2006-01-02"`
}

type Salary struct {
	EmpNo    int
	Salary   int
	FromDate time.Time `mgorm:"layout=2006-01-02"`
	ToDate   time.Time `mgorm:"layout=2006-01-02"`
}

type Title struct {
	EmpNo    int
	Title    string
	FromDate time.Time `mgorm:"layout=2006-01-02"`
	ToDate   time.Time `mgorm:"layout=2006-01-02"`
}
