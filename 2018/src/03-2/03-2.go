package main

import (
	"fmt"
	"regexp"
	"au"
)

type Claim struct {
	id string
	x int
	y int
	width int
	height int
}

func readClaims() ([]Claim) {
	re := regexp.MustCompile("\\#([0-9]*) @ ([0-9]*),([0-9]*): ([0-9]*)x([0-9]*)");

	var claims []Claim;
	for _, input := range au.ReadInputAsStringArray("03") {
		matches := re.FindStringSubmatch(input);

		id := matches[1];
		x := au.ToNumber(matches[2]);
		y := au.ToNumber(matches[3]);
		width := au.ToNumber(matches[4]);
		height := au.ToNumber(matches[5]);

		claims = append(claims, Claim{id, x, y, width, height})
	}

	return claims;
}

func testClaims() ([]Claim) {
	return []Claim {
		Claim{"1", 1, 3, 4, 4},
		Claim{"2", 3, 1, 4, 4},
		Claim{"3", 5, 5, 2, 2},
	}
}

type Fabric struct {
	squares [1000][1000][]string
}

func (this *Fabric) mark(claim Claim) {
	for dx := 0; dx < claim.width; dx++ {
		for dy := 0; dy < claim.height; dy++ {
			this.squares[claim.x + dx][claim.y + dy] = append(this.squares[claim.x + dx][claim.y + dy], claim.id);
		}
	}
}

func (this *Fabric) getOverlappingClaims() map[string] bool {
	overlappingClaims := map[string] bool{}

	for _, row := range this.squares {
		for _, square := range row {
			if len(square) > 1 {
				for _, claimId := range square {
					overlappingClaims[claimId] = true;
				}
			}
		}
	}

	return overlappingClaims
}

func main() {
	claims := readClaims()
	// claims := testClaims();

	fabric := Fabric{}

	for _, claim := range claims {
		fabric.mark(claim)
	}

	overlappingClaims := fabric.getOverlappingClaims()
	for _, claim := range claims {
		_, ok := overlappingClaims[claim.id]
		if (!ok) {
			fmt.Println(claim.id)
		}
	}
}
