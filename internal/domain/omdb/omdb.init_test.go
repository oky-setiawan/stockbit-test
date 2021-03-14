package omdb

import (
	"github.com/oky-setiawan/stockbit-test/internal/config"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInit(t *testing.T) {
	g := NewGomegaWithT(t)
	cfg := &config.PartnerConfig{OMDB: config.OMDBConfig{}}
	m := Init(cfg)
	g.Expect(m).ShouldNot(BeNil())
}
