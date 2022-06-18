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

	type testCase struct {
		ct   timeutils.ClockTime
		want clockface.Line
	}
	cases := []testCase{
		{
			timeutils.ClockTime{H: 0, M: 0, S: 0},
			clockface.Line{150, 150, 150, 60},
		},
		{
			timeutils.ClockTime{H: 0, M: 0, S: 30},
			clockface.Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(c.ct.String(), func(t *testing.T) {
			tm := c.ct.ToTime()
			b := bytes.Buffer{}
			clockface.New(origin, scale).WriteSVG(&b, tm)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !isIn(c.want, svg.Lines) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.want, svg.Lines)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	origin := 150.0
	scale := 90.0

	type testCase struct {
		ct   timeutils.ClockTime
		want clockface.Line
	}
	cases := []testCase{
		{
			timeutils.ClockTime{H: 0, M: 0, S: 0},
			clockface.Line{150, 150, 150, 60},
		},
		{
			timeutils.ClockTime{H: 0, M: 30, S: 0},
			clockface.Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(c.ct.String(), func(t *testing.T) {
			tm := c.ct.ToTime()
			b := bytes.Buffer{}
			clockface.New(origin, scale).WriteSVG(&b, tm)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !isIn(c.want, svg.Lines) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.want, svg.Lines)
			}
		})
	}
}

func isIn[T comparable](l T, ls []T) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}

	return false

}
