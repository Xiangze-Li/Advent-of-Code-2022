package puzzle23

import (
	"advent2022/base"
	"advent2022/util"
	"maps"
	"math"
)

type p struct {
	elfs map[[2]int]bool
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.elfs = make(map[[2]int]bool)

	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				p.elfs[[2]int{i, j}] = true
			}
		}
	}
}

func moveRound(round int, elfs *map[[2]int]bool) bool {
	moves := map[[2]int][2]int{}
	proposeCount := map[[2]int]int{}

	for elf := range *elfs {
		adj := util.Associative(allDir[:], func(dir direction) (direction, bool) {
			d := delta[dir]
			return dir, (*elfs)[[2]int{elf[0] + d[0], elf[1] + d[1]}]
		})
		if !util.ReduceMap(adj, func(b bool, _ direction, exist bool) bool { return b || exist }, false) {
			continue
		}

		canMove := false
		var dir direction
		for k := round; k < round+4; k++ {
			canMoveDir := true
			for _, d := range checkDirRange[checkDir[k%4]] {
				if adj[d] {
					canMoveDir = false
					break
				}
			}
			if canMoveDir {
				canMove = true
				dir = checkDir[k%4]
				break
			}
		}
		if canMove {
			d := delta[dir]
			moves[elf] = [2]int{elf[0] + d[0], elf[1] + d[1]}
			proposeCount[moves[elf]]++
		}
	}

	if len(moves) == 0 {
		return false
	}
	for from, to := range moves {
		if proposeCount[to] > 1 {
			continue
		}
		delete(*elfs, from)
		(*elfs)[to] = true
	}
	return true
}

func (p *p) Solve1() any {
	elfs := maps.Clone(p.elfs)

	for round := 0; round < 10; round++ {
		_ = moveRound(round, &elfs)
	}

	minI, maxI := math.MaxInt, math.MinInt
	minJ, maxJ := math.MaxInt, math.MinInt
	for elf := range elfs {
		minI = min(minI, elf[0])
		maxI = max(maxI, elf[0])
		minJ = min(minJ, elf[1])
		maxJ = max(maxJ, elf[1])
	}

	return (maxI-minI+1)*(maxJ-minJ+1) - len(elfs)
}

func (p *p) Solve2() any {
	elfs := maps.Clone(p.elfs)

	round := 0
	for moveRound(round, &elfs) {
		round++
	}

	return round + 1
}

func init() {
	base.Register(23, &p{})
}

type direction byte

const (
	N direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

var checkDir = [...]direction{N, S, W, E}
var allDir = [...]direction{N, NE, E, SE, S, SW, W, NW}

var checkDirRange = map[direction][3]direction{
	N: {NW, N, NE},
	S: {SW, S, SE},
	W: {NW, W, SW},
	E: {NE, E, SE},
}

var delta = map[direction][2]int{
	N:  {-1, 0},
	S:  {1, 0},
	W:  {0, -1},
	E:  {0, 1},
	NE: {-1, 1},
	NW: {-1, -1},
	SE: {1, 1},
	SW: {1, -1},
}
