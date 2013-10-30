package char

import (
	"testing"

	"github.com/knakk/specs"
)

func TestRemovePunctuationFilter(t *testing.T) {
	s := specs.New(t)

	charf := NewRemovePunctuationFilter()

	s.Expect(charf.Filter("hei, `du`."), "hei du")
	s.Expect(charf.Filter("hjælp!! `skrekk-blandet*?=) \"fryd\""),
		"hjælp skrekkblandet fryd")
}

func TestPunct2SpaceFilter(t *testing.T) {
	s := specs.New(t)

	charf := NewPunct2SpaceFilter()

	s.Expect(charf.Filter("Skrap-balle,kan.du\"fikse\"meg?vennligst~svar."),
		"Skrap balle kan du fikse meg vennligst svar ")
}

func TestNewNorwegianFoldingFilter(t *testing.T) {
	s := specs.New(t)

	charf := NewNorwegianFoldingFilter()

	s.Expect(charf.Filter("éntrö Él ñinjä misćjø bånd"),
		"entrø El ninjæ miscjø bånd")
}
