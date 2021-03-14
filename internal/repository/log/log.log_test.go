package log

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	. "github.com/onsi/gomega"
	"testing"
)

func Test_logRepository_LogAction(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		doRepoTest(t, func(g *GomegaWithT, m repoTestObject) {
			req := &entity.LogActionRequest{}

			m.dbMock.ExpectExec("INSERT INTO log").WillReturnResult(sqlmock.NewResult(1, 1))

			err := m.repo.LogAction(m.ctx, req)
			g.Expect(err).Should(BeNil())
		})
	})
}
