package dice

import (
	"math/rand"
	"sort"
	"time"
)

// Roller store information regarding a dice roll.
type Roller struct {
	rng       rand.Rand
	dice      []Dice
	modifiers map[string]int
	dropLow   int
	dropHigh  int
	keepHigh  int
	keepLow   int
	fixed     bool
	fix       int
}

// Dice store information regarding dice.
type Dice struct {
	amount int
	faces  int
}

// New create a new empty roller.
func New() *Roller {
	return &Roller{
		rng:       rand.Rand{},
		dice:      []Dice{},
		modifiers: map[string]int{},
		dropLow:   0,
		dropHigh:  0,
		keepHigh:  0,
		keepLow:   0,
	}
}

// Source set the random source for the roller.
func (r *Roller) Rng(rng rand.Rand) *Roller {
	r.rng = rng
	return r
}

// Dice add dice to the Roller.
func (r *Roller) Dice(amount, faces int) *Roller {
	r.dice = append(r.dice, Dice{
		amount: amount,
		faces:  faces,
	})
	return r
}

// AddModifier add modifier with given name to roller.
func (r *Roller) AddModifier(name string, value int) *Roller {
	r.modifiers[name] = value
	return r
}

// RemoveModifer remove modifier with given name from roller.
func (r *Roller) RemoveModifier(name string) *Roller {
	delete(r.modifiers, name)
	return r
}

// DropLow set amount of lowest dice to be removed from the roll.
func (r *Roller) DropLow(value int) *Roller {
	r.dropLow = value
	return r
}

// DropHigh set amount of highest dice to be removed from the roll.
func (r *Roller) DropHigh(value int) *Roller {
	r.dropHigh = value
	return r
}

// KeepHigh set amount of highest dice to be kept in the roll.
func (r *Roller) KeepHigh(value int) *Roller {
	r.keepHigh = value
	return r
}

// KeepLow set amount of lowest dice to be kept in the roll.
func (r *Roller) KeepLow(value int) *Roller {
	r.keepLow = value
	return r
}

// Fix the roller to output a certain value
func (r *Roller) Fix(value int) *Roller {
	r.fixed = true
	r.fix = value
	return r
}

// Remove fix from the roller, has no effect if no fix is set
func (r *Roller) Unfix() *Roller {
	r.fixed = false
	r.fix = 0
	return r
}

// Roll roll dice, apply operations, and add modifiers returning the total.
func (r *Roller) Roll() int {
	if r.fixed {
		return r.fix
	}

	results := []int{}

	if r.rng == (rand.Rand{}) {
		r.rng = *rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	for _, dice := range r.dice {
		for i := 0; i < dice.amount; i++ {
			results = append(results, r.rng.Intn(dice.faces)+1)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return i > j
	})

	if r.dropHigh > 0 {
		results = results[:len(results)-r.dropHigh]
	}
	if r.dropLow > 0 {
		results = results[r.dropLow:]
	}
	if r.keepHigh > 0 {
		results = results[len(results)-r.keepHigh:]
	}
	if r.keepLow > 0 {
		results = results[:r.keepLow]
	}

	total := 0
	for _, result := range results {
		total += result
	}

	for _, value := range r.modifiers {
		total += value
	}

	return total
}
