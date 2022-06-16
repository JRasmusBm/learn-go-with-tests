package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/JRasmusBm/learn-go-with-tests/pkg/clockface"
)

func TestSVGWriter(t *testing.T) {
	origin := 150.0
	scale := 90.0

	t.Run("- After midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		b := bytes.Buffer{}
		clockface.New(origin, scale).WriteSVG(&b, tm)

		svg := clockface.SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		want := clockface.Line{150, 150, 150, 60}

		if !isInLines(want, svg.Lines) {
			t.Errorf("Expected to find the second hand %+v, in the SVG output %v", want, b.String())
		}

	})
}

func isInLines(l clockface.Line, ls []clockface.Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}

	return false

}
