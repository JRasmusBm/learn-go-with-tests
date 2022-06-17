package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/JRasmusBm/learn-go-with-tests/pkg/clockface"
	"github.com/JRasmusBm/learn-go-with-tests/pkg/timeutils"
)

func TestSVGWriterSecondHand(t *testing.T) {
	origin := 150.0
	scale := 90.0

	t.Run("- After midnight", func(t *testing.T) {
		tm := timeutils.ClockTime{H: 0, M: 0, S: 0}.ToTime()
		b := bytes.Buffer{}
		clockface.New(origin, scale).WriteSVG(&b, tm)

		svg := clockface.SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		want := clockface.Line{150, 150, 150, 60}

		if !isIn(want, svg.Lines) {
			t.Errorf("Expected to find the second hand %+v, in the SVG output %v", want, b.String())
		}
	})
}

func isIn[T comparable](l T, ls []T) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}

	return false

}
