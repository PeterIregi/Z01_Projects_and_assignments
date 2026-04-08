package ascii

import "testing"
import "ascii-art-web/asciiart"

func TestReader(t *testing.T) {
	lines, _ := asciiarts.ReadLinesFromFile("../banners/standard.txt")
	got := len(lines)

	want := 855

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
