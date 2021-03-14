package log

import (
	"context"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/oky-setiawan/stockbit-test/lib/database"
	. "github.com/onsi/gomega"
	"testing"
)

func TestNew(t *testing.T) {
	g := NewGomegaWithT(t)
	m := New(&Opts{})
	g.Expect(m).ShouldNot(BeNil())
	g.Expect(m).Should(BeAssignableToTypeOf(&logRepository{}))
}

type repoTestObject struct {
	ctx    context.Context
	dbMock sqlmock.Sqlmock
	repo   *logRepository
}

func doRepoTest(t *testing.T, fn func(g *GomegaWithT, m repoTestObject)) {
	g := NewGomegaWithT(t)

	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed do repo test")
	}

	sqlDB := sqlx.NewDb(db, "sqlmock")
	mockDBStore := &database.Store{
		Master: sqlDB,
		Slave:  sqlDB,
	}

	repo := &logRepository{
		db: mockDBStore,
	}

	m := repoTestObject{
		ctx:    context.Background(),
		dbMock: dbMock,
		repo:   repo,
	}
	fn(g, m)
}
