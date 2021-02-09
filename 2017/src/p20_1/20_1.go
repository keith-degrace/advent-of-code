package p20_1

import (
	"au"
	"fmt"
	"regexp"
)

type Particle struct {
	manhattanPosition     int
	manhattanVelocity     int
	manhattanAcceleration int
}

func (p *Particle) isCloserToOriginThan(other Particle) bool {
	if p.manhattanAcceleration != other.manhattanAcceleration {
		return p.manhattanAcceleration < other.manhattanAcceleration
	}

	if p.manhattanVelocity != other.manhattanVelocity {
		return p.manhattanVelocity < other.manhattanVelocity
	}

	return p.manhattanPosition < other.manhattanPosition
}

func getDistanceFromOrigin(x, y, z int) int {
	return au.AbsInt(x) + au.AbsInt(y) + au.AbsInt(z)
}

func load(input []string) []Particle {
	particles := []Particle{}

	re := regexp.MustCompile("p=<(.*),(.*),(.*)>, v=<(.*),(.*),(.*)>, a=<(.*),(.*),(.*)>")
	for _, line := range input {
		match := re.FindStringSubmatch(line)

		particle := Particle{}

		particle.manhattanPosition = getDistanceFromOrigin(
			au.ToNumber(match[1]),
			au.ToNumber(match[2]),
			au.ToNumber(match[3]))

		particle.manhattanVelocity = getDistanceFromOrigin(
			au.ToNumber(match[4]),
			au.ToNumber(match[5]),
			au.ToNumber(match[6]))

		particle.manhattanAcceleration = getDistanceFromOrigin(
			au.ToNumber(match[7]),
			au.ToNumber(match[8]),
			au.ToNumber(match[9]))

		particles = append(particles, particle)
	}

	return particles
}

func Solve() {
	input := au.ReadInputAsStringArray("20")

	particles := load(input)

	closestToOrigin := 0

	for i := 1; i < len(particles); i++ {
		if particles[i].isCloserToOriginThan(particles[closestToOrigin]) {
			closestToOrigin = i
		}
	}

	fmt.Println(closestToOrigin)
}
