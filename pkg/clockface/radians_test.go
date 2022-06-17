package clockface

import (
	"math"
	"reflect"
	"testing"

	"github.com/JRasmusBm/learn-go-with-tests/pkg/timeutils"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		ct   timeutils.ClockTime
		want float64
	}{
		{ct: timeutils.ClockTime{H: 0, M: 0, S: 0}, want: 0.0},
		{ct: timeutils.ClockTime{H: 0, M: 0, S: 30}, want: math.Pi},
		{ct: timeutils.ClockTime{H: 0, M: 0, S: 45}, want: (math.Pi / 2) * 3},
		{ct: timeutils.ClockTime{H: 0, M: 0, S: 7}, want: (math.Pi / 30) * 7},
	}
	for _, c := range cases {
		t.Run(c.ct.String(), func(t *testing.T) {
			got := secondsInRadians(c.ct.ToTime())
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%v, got %v want %v", c.ct, got, c.want)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	t.Run("ClockTimes", func(t *testing.T) {
		type testCase struct {
			name string
			ct   timeutils.ClockTime
			want Point
		}

		cases := []testCase{
			{ct: timeutils.ClockTime{H: 0, M: 0, S: 0}, want: Point{0, 1}},
			{ct: timeutils.ClockTime{H: 0, M: 0, S: 15}, want: Point{1, 0}},
			{ct: timeutils.ClockTime{H: 0, M: 0, S: 30}, want: Point{0, -1}},
			{ct: timeutils.ClockTime{H: 0, M: 0, S: 45}, want: Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(c.ct.String(), func(t *testing.T) {
				got := secondHandPoint(c.ct.ToTime())

				if !got.Equals(c.want) {
					t.Errorf("%#v, got %v want %v", c.ct, got, c.want)
				}
			})
		}
	})

}
