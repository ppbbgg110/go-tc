package core

import (
	"testing"
)

func TestClock(t *testing.T) {
	t.Run("Tick", func(t *testing.T) {
		tick := Time2Tick(0xC0FFEE)
		time := Tick2Time(tick)
		// to avoid the dependency to math and make sure, diff is always equal or
		// bigger than 0, diff is the square of the difference of both values
		diff := (0xC0FFEE - time) * (0xC0FFEE - time)
		if diff > 3 {
			t.Fatalf("expected %d, got %d", 0xC0FFEE, time)
		}
	})

	t.Run("Ktime", func(t *testing.T) {
		ktime := Time2Ktime(0xC0FFEE)
		time := Ktime2Time(ktime)
		diff := (0xC0FFEE - time) * (0xC0FFEE - time)
		if diff > 3 {
			t.Fatalf("expected %d, got %d", 0xC0FFEE, time)
		}
	})

	t.Run("XmitTime", func(t *testing.T) {
		timeA := XmitTime(4096, 4096)
		timeB := Time2Tick(timeUnitsPerSec)
		if timeA != timeB {
			t.Fatalf("expected %d, got %d", timeB, timeA)
		}
	})
	t.Run("XmitSize", func(t *testing.T) {
		a := XmitSize(timeUnitsPerSec, 1)
		b := Tick2Time(1)
		if a != b {
			t.Fatalf("expected %d, got %d", a, b)
		}
	})
}
