package main_test

import (
	dice "dice"
	"testing"
)

// TestEmptyRoller create an empty roller,
// checking that a roll returns 0.
func TestEmptyRoller(t *testing.T) {
	result := dice.New().Roll()
	if result != 0 {
		t.Fatalf(`NewROller().Roll() = %d, should be = 0`, result)
	}
}

// Test3d6 create a roller with a 3d6, checking
// that 1000 rolls are between 3 and 18.
func Test3d6(t *testing.T) {
	for i := 0; i < 1000; i++ {
		result := dice.New().Dice(3, 6).Roll()
		if result < 3 {
			t.Fatalf(`NewRoller().Dice(3, 6).Roll() = %d, should be >= 3`, result)
		}

		if result > 3*6 {
			t.Fatalf(`NewRoller().Dice(3, 6).Roll() = %d, should be <= 18`, result)
		}
	}
}

// Test4d6DropLowest create a roller of 4d6dl1,
// checking that 1000 rolls are between 3 and 18.
func Test4d6DropLowest(t *testing.T) {
	for i := 0; i < 1000; i++ {
		result := dice.New().Dice(4, 6).DropLow(1).Roll()
		if result < 3 {
			t.Fatalf(`NewRoller().Dice(4, 6).DropLow(1).Roll() = %d, should be >= 3`, result)
		}
		if result > 3*6 {
			t.Fatalf(`NewRoller().Dice(4, 6).DropLow(1).Roll() = %d, should be <= 18`, result)
		}
	}
}

// Test4d6DropHighest create a roller of 4d6dh1,
// checking that 1000 rolls are between 3 and 18.
func Test4d6DropHighest(t *testing.T) {
	for i := 0; i < 1000; i++ {
		result := dice.New().Dice(4, 6).DropLow(1).Roll()
		if result < 3 {
			t.Fatalf(`NewRoller().Dice(4, 6).DropHigh(1).Roll() = %d, should be >= 3`, result)
		}
		if result > 3*6 {
			t.Fatalf(`NewRoller().Dice(4, 6).DropHigh(1).Roll() = %d, should be <= 18`, result)
		}
	}
}

// TestNestedRoll create a roller of 1d(1d2),
// checking that 1000 rolls are between 1 and 2.
func TestNestedRoll(t *testing.T) {
	inner := dice.New().Dice(1, 2)
	for i := 0; i < 1000; i++ {
		result := dice.New().Dice(1, inner.Roll()).Roll()
		if result < 1 {
			t.Fatalf(`NewRoller().Dice(1, NewRoller().Dice(1, 2).Roll()).Roll() = %d, should be >= 1`, result)
		}
		if result > 2 {
			t.Fatalf(`NewRoller().Dice(1, NewRoller().Dice(1, 2).Roll()).Roll() = %d, should be <= 2`, result)
		}
	}
}
