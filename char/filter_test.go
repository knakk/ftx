package char

import (
	"testing"
)

func TestRemovePunctuationFilter(t *testing.T) {
	charf := NewRemovePunctuationFilter()
	input := "hei, `du`."
	want := "hei du"
	if charf.Filter(input) != want {
		t.Errorf("RemovePunctuationFilter(%q) => %v; want %v", input, charf.Filter(input), want)
	}
	input = "hjælp!! `skrekk-blandet*?=) \"fryd\""
	want = "hjælp skrekkblandet fryd"
	if charf.Filter(input) != want {
		t.Errorf("RemovePunctuationFilter(%q) => %v; want %v", input, charf.Filter(input), want)
	}
}

func TestPunct2SpaceFilter(t *testing.T) {
	charf := NewPunct2SpaceFilter()
	input := "Skrap-balle,kan.du\"fikse\"meg?vennligst~svar."
	want := "Skrap balle kan du fikse meg vennligst svar "
	if charf.Filter(input) != want {
		t.Errorf("Punct2SpaceFilter(%q) => %v; want %v", input, charf.Filter(input), want)
	}
}

func TestNewNorwegianFoldingFilter(t *testing.T) {
	charf := NewNorwegianFoldingFilter()
	input := "éntrö Él ñinjä misćjø bånd"
	want := "entrø El ninjæ miscjø bånd"
	if charf.Filter(input) != want {
		t.Errorf("NorwegianFoldingFilter(%q) => %v; want %v", input, charf.Filter(input), want)
	}
}
