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
				got := New(0, 1).SecondHand(c.ct.ToTime())
				got.Y = -got.Y

				if !got.Equals(c.want) {
					t.Errorf("%#v, got %v want %v", c.ct, got, c.want)
				}
			})
		}
	})
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		ct   timeutils.ClockTime
		want float64
	}{
		{ct: timeutils.ClockTime{H: 0, M: 0, S: 0}, want: 0.0},
		{ct: timeutils.ClockTime{H: 0, M: 0, S: 7}, want: 7 * (math.Pi / (30 * 60))},
	}
	for _, c := range cases {
		t.Run(c.ct.String(), func(t *testing.T) {
			got := minutesInRadians(c.ct.ToTime())
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%v, got %v want %v", c.ct, got, c.want)
			}
		})
	}
}

func TestMinuteHandVector(t *testing.T) {
	t.Run("ClockTimes", func(t *testing.T) {
		type testCase struct {
			name string
			ct   timeutils.ClockTime
			want Point
		}

		cases := []testCase{
			{ct: timeutils.ClockTime{H: 0, M: 0, S: 0}, want: Point{0, 1}},
			{ct: timeutils.ClockTime{H: 0, M: 15, S: 0}, want: Point{1, 0}},
			{ct: timeutils.ClockTime{H: 0, M: 30, S: 0}, want: Point{0, -1}},
			{ct: timeutils.ClockTime{H: 0, M: 45, S: 0}, want: Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(c.ct.String(), func(t *testing.T) {
				got := New(0, 1).MinuteHand(c.ct.ToTime())
				got.Y = -got.Y

				if !got.Equals(c.want) {
					t.Errorf("%#v, got %v want %v", c.ct, got, c.want)
				}
			})
		}
	})
}
func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		ct   timeutils.ClockTime
		want float64
	}{
		{ct: timeutils.ClockTime{H: 0, M: 0, S: 0}, want: 0},
		{ct: timeutils.ClockTime{H: 6, M: 0, S: 0}, want: math.Pi},
		{ct: timeutils.ClockTime{H: 21, M: 0, S: 0}, want: math.Pi * 1.5},
	}
	for _, c := range cases {
		t.Run(c.ct.String(), func(t *testing.T) {
			got := hoursInRadians(c.ct.ToTime())
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%v, got %v want %v", c.ct, got, c.want)
			}
		})
	}
}
