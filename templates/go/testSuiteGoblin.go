package kataname

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestKataName(t *testing.T) {
	g := Goblin(t)

	g.Describe("kataName", func() {
		g.It("example test", func() {
			g.Assert("example").Equal(kataName())
		})

	})
}
