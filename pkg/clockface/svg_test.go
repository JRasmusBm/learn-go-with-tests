package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/JRasmusBm/learn-go-with-tests/pkg/clockface"
)

func TestSvgWriter(t *testing.T) {
	origin := 150.0
	scale := 90.0

	t.Run("- After midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		b := bytes.Buffer{}
		clockface.New(origin, scale).WriteSvg(&b, tm)

		svg := clockface.SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		x2 := "150.000"
		y2 := "60.000"

		for _, line := range svg.Line {
			if line.X2 == x2 && line.Y2 == y2 {
				return
			}
		}

		t.Errorf("Expected to find the second hand with x2 of %+v and y2 of %+v, in the SVG output %v", x2, y2, b.String())
	})
}
