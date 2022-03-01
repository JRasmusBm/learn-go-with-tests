package clockface

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"
)

type clockTime struct {
	h int
	m int
	s int
}

func (c clockTime) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", c.h, c.m, c.s)
}

func (ct clockTime) toTime() time.Time {
	return time.Date(312, time.October, 28, ct.h, ct.m, ct.s, 0, time.UTC)
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		ct   clockTime
		want float64
	}{
		{ct: clockTime{h: 0, m: 0, s: 0}, want: 0.0},
		{ct: clockTime{h: 0, m: 0, s: 30}, want: math.Pi},
		{ct: clockTime{h: 0, m: 0, s: 45}, want: (math.Pi / 2) * 3},
		{ct: clockTime{h: 0, m: 0, s: 7}, want: (math.Pi / 30) * 7},
	}
	for _, c := range cases {
		t.Run(c.ct.String(), func(t *testing.T) {
			got := secondsInRadians(c.ct.toTime())
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
			ct   clockTime
			want Point
		}

		cases := []testCase{
			{ct: clockTime{h: 0, m: 0, s: 0}, want: Point{0, 1}},
			{ct: clockTime{h: 0, m: 0, s: 15}, want: Point{1, 0}},
			{ct: clockTime{h: 0, m: 0, s: 30}, want: Point{0, -1}},
			{ct: clockTime{h: 0, m: 0, s: 45}, want: Point{-1, -1}},
		}

		for _, c := range cases {
			t.Run(c.ct.String(), func(t *testing.T) {
				got := secondHandPoint(c.ct.toTime())

				if !got.Equals(c.want) {
					t.Errorf("%#v, got %v want %v", c.ct, got, c.want)
				}
			})
		}
	})

}
