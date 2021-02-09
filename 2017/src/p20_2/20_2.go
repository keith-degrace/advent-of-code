package p20_2

import (
	"au"
	"fmt"
	"regexp"
)

func getDistanceFromOrigin(x, y, z int) int {
	return au.AbsInt(x) + au.AbsInt(y) + au.AbsInt(z)
}

type Particle struct {
	p    []int
	v    []int
	a    []int
	dead bool
}

func (p *Particle) willNeverCollideOnAxis(o *Particle, axis int) bool {
	isNegativeVelocityP := p.v[axis] < 0
	isNegativeAccelerationP := p.a[axis] < 0

	isNegativeVelocityO := o.v[axis] < 0
	isNegativeAccelerationO := o.a[axis] < 0

	// If one of the points is not moving and accelerating in the same direction, we can't easily know if they will ever collide.
	if isNegativeVelocityP != isNegativeAccelerationP || isNegativeVelocityO != isNegativeAccelerationO {
		return false
	}

	return (p.a[axis] < o.a[axis]) == (p.v[axis] < o.v[axis]) == (p.p[axis] < o.p[axis])
}

func (p *Particle) willNeverCollideWith(o *Particle) bool {
	return p.willNeverCollideOnAxis(o, 0) || p.willNeverCollideOnAxis(o, 1) || p.willNeverCollideOnAxis(o, 2)
}

func (p *Particle) collidesWith(o *Particle) bool {
	return p.p[0] == o.p[0] && p.p[1] == o.p[1] && p.p[2] == o.p[2]
}

func load(input []string) []*Particle {
	particles := []*Particle{}

	re := regexp.MustCompile("p=<(.*),(.*),(.*)>, v=<(.*),(.*),(.*)>, a=<(.*),(.*),(.*)>")
	for _, line := range input {
		match := re.FindStringSubmatch(line)

		particle := new(Particle)
		particle.dead = false
		particle.p = []int{
			au.ToNumber(match[1]),
			au.ToNumber(match[2]),
			au.ToNumber(match[3])}

		particle.v = []int{
			au.ToNumber(match[4]),
			au.ToNumber(match[5]),
			au.ToNumber(match[6])}

		particle.a = []int{
			au.ToNumber(match[7]),
			au.ToNumber(match[8]),
			au.ToNumber(match[9])}

		particles = append(particles, particle)
	}

	return particles
}

func tick(particles []*Particle) {
	for _, particle := range particles {
		if particle.dead {
			continue
		}

		particle.v[0] += particle.a[0]
		particle.v[1] += particle.a[1]
		particle.v[2] += particle.a[2]

		particle.p[0] += particle.v[0]
		particle.p[1] += particle.v[1]
		particle.p[2] += particle.v[2]
	}
}

func Solve() {
	input := au.ReadInputAsStringArray("20")

	particles := load(input)

	for i := 0; i < 1000; i++ {

		// Check for collisions
		for i := 0; i < len(input)-1; i++ {
			if particles[i].dead {
				continue
			}

			for j := i + 1; j < len(input); j++ {
				if particles[j].dead {
					continue
				}

				if particles[i].collidesWith(particles[j]) {
					particles[i].dead = true
					particles[j].dead = true
				}
			}
		}

		tick(particles)

		// If there are no more potentially colliding particles, we are done...

		// Cheesed out here.  I was trying to detect when there are no longer any potentially colliding
		// pairs (and it's not working yet) but it turns out that 1000 iterations gave us a gold star.

		// {
		// 	potentiallyCollidingPairs := [][]*Particle{}

		// 	for i := 0; i < len(input)-1; i++ {
		// 		if particles[i].dead {
		// 			continue
		// 		}

		// 		for j := i + 1; j < len(input); j++ {
		// 			if particles[j].dead {
		// 				continue
		// 			}

		// 			if particles[i].collidesWith(particles[j]) {
		// 				particles[i].dead = true
		// 				particles[j].dead = true
		// 			} else if !particles[i].willNeverCollideWith(particles[j]) {
		// 				potentiallyCollidingPairs = append(potentiallyCollidingPairs, []*Particle{particles[i], particles[j]})
		// 			}
		// 		}
		// 	}

		// 	if len(potentiallyCollidingPairs) == 0 {
		// 		fmt.Println("")
		// 		break
		// 	}
		// }
	}

	count := 0
	for _, particle := range particles {
		if !particle.dead {
			count++
		}
	}

	fmt.Println(count)
}
