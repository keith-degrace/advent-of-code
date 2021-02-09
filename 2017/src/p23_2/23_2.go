package p23_2

import (
	"fmt"
)

func runAsCode() int {
	a := 1
	b := 57
	c := b
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0

	if a != 0 {
		b *= 100
		b += 100000
		c = b
		c += 17000
	}

	for {
		f = 1
		d = 2

		for f == 1 {
			e = 2

			for f == 1 {
				g = d
				g *= e
				g -= b
				if g == 0 {
					f = 0
				}

				// *************************
				// This is the key. There are two nested both iterating from 2 to the current value of 'b', and it tries to find
				// a pair of numbers that when multiplied, equals to b.  We can speed this up by stopping the loop whenever the
				// number goes beyond b...
				// *************************
				if g > 0 {
					break
				}

				e += 1
				g = e
				g -= b

				if g == 0 {
					break
				}
			}

			d += 1
			g = d
			g -= b

			if g == 0 {
				break
			}
		}

		if f == 0 {
			h += 1
		}

		g = b
		g -= c
		b += 17

		if g == 0 {
			break
		}
	}

	return h
}

func Solve() {
	// Can't say I'm a fan of these.  Pretty much had to analyze the code
	// and figure out what it was doing, and optimize it so I converted it
	// to GO code and messed around with it.
	fmt.Println(runAsCode())
}
