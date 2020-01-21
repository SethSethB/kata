package cmd

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestHelpersSuite(t *testing.T) {
	g := Goblin(t)

	g.Describe("ConvertToCamelCase", func() {
		g.It("Handles simple case", func() {
			g.Assert("simple").Equal(convertToCamelCase("simple"))
		})

		g.It("Converts words separated by space", func() {
			g.Assert("lessSimple").Equal(convertToCamelCase("less Simple"))
		})

		g.It("Converts case mixtures", func() {
			g.Assert("caseMixtureName").Equal(convertToCamelCase("CasE mIXture namE"))
		})

		g.It("Handles multiple spaces", func() {
			g.Assert("multipleSpaces").Equal(convertToCamelCase("MultiPle    spaceS"))
		})

	})
}
