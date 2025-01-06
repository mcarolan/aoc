package util

import "testing"

func TestDirectionClockwise(t *testing.T) {
	if North.Clockwise() != East {
		t.Error("North clockwise")
	}
	if East.Clockwise() != South {
		t.Error("East clockwise")
	}
	if South.Clockwise() != West {
		t.Error("South clockwise")
	}
	if West.Clockwise() != North {
		t.Error("West clockwise")
	}
}

func TestDirectionAntiClockwise(t *testing.T) {
	if North.AntiClockwise() != West {
		t.Errorf("North anticlockwise %d", North.AntiClockwise())
	}
	if East.AntiClockwise() != North {
		t.Error("East anticlockwise")
	}
	if South.AntiClockwise() != East {
		t.Error("South anticlockwise")
	}
	if West.AntiClockwise() != South {
		t.Error("West anticlockwise")
	}
}
