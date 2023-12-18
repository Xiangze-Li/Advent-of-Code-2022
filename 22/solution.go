package puzzle22

import (
	"advent2022/base"
	"advent2022/util"
	"regexp"
)

type gridType uint8

const (
	void gridType = iota
	open
	wall
)

type p struct {
	grids [][]gridType
	dists []int64
	turns []bool

	startPos [2]int

	lenI, lenJ int
}

func (p *p) inbound(i, j int) bool {
	return i >= 0 && i < p.lenI && j >= 0 && j < p.lenJ
}

func (p *p) Init(filename string) {
	blocks := util.GetBlocks(filename)

	lenI := len(blocks[0])
	lenJ := util.Reduce(blocks[0], func(maxium int, line string) int {
		return max(maxium, len(line))
	}, 0)

	p.grids = util.SliceND[gridType](lenI, lenJ).([][]gridType)

	for i, line := range blocks[0] {
		for j, c := range line {
			switch c {
			case '.':
				p.grids[i][j] = open
			case '#':
				p.grids[i][j] = wall
			}
		}
	}

	for j, c := range p.grids[0] {
		if c == open {
			p.startPos = [2]int{0, j}
			break
		}
	}

	reDist := regexp.MustCompile(`\d+`)
	reTurn := regexp.MustCompile(`L|R`)
	p.dists = util.ArrayStrToInt64(reDist.FindAllString(blocks[1][0], -1))
	turn := reTurn.FindAllString(blocks[1][0], -1)
	p.turns = make([]bool, len(turn))
	for i, t := range turn {
		p.turns[i] = t == "L"
	}

	util.Assert(len(p.dists) == len(p.turns)+1, "len(p.dists) != len(p.turns)+1")

	p.lenI = lenI
	p.lenJ = lenJ
}

func (p *p) Solve1() any {
	cur := p.startPos
	dir := E

	for pc, nStep := range p.dists {
		for step := int64(0); step < nStep; step++ {
			next := [2]int{cur[0] + deltas[dir][0], cur[1] + deltas[dir][1]}
			if !p.inbound(next[0], next[1]) || p.grids[next[0]][next[1]] == void {
				switch dir {
				case E:
					var j int
					for j = cur[1]; j >= 0; j-- {
						if p.grids[cur[0]][j] == void {
							break
						}
					}
					next[1] = j + 1
				case S:
					var i int
					for i = cur[0]; i >= 0; i-- {
						if p.grids[i][cur[1]] == void {
							break
						}
					}
					next[0] = i + 1
				case W:
					var j int
					for j = cur[1]; j < p.lenJ; j++ {
						if p.grids[cur[0]][j] == void {
							break
						}
					}
					next[1] = j - 1
				case N:
					var i int
					for i = cur[0]; i < p.lenI; i++ {
						if p.grids[i][cur[1]] == void {
							break
						}
					}
					next[0] = i - 1
				}
			}
			if p.grids[next[0]][next[1]] == wall {
				break
			}
			cur = next
		}
		if pc != len(p.dists)-1 {
			dir = dir.turn(p.turns[pc])
		}
	}

	return 1000*(cur[0]+1) + 4*(cur[1]+1) + int(dir)
}

// wrap2 wraps coordinate off-board to correct location, together with new direction
func (p *p) wrap2(next [2]int) ([2]int, direction) {
	const side = 50

	switch {
	case next[0] == 0*side-1 && 1*side <= next[1] && next[1] < 2*side:
		delta := next[1] - 1*side
		return [2]int{3*side + delta, 0}, E
	case next[0] == 0*side-1 && 2*side <= next[1] && next[1] < 3*side:
		delta := next[1] - 2*side
		return [2]int{4*side - 1, 0 + delta}, N
	case next[1] == 1*side-1 && 0*side <= next[0] && next[0] < 1*side:
		delta := next[0] - 0*side
		return [2]int{3*side - 1 - delta, 0}, E
	case next[1] == 1*side-1 && 1*side <= next[0] && next[0] < 2*side:
		delta := next[0] - 1*side
		return [2]int{2 * side, 0*side + delta}, S
	case next[1] == 3*side && 0*side <= next[0] && next[0] < 1*side:
		delta := next[0] - 0*side
		return [2]int{3*side - 1 - delta, 2*side - 1}, W
	case next[0] == 1*side && 2*side <= next[1] && next[1] < 3*side:
		delta := next[1] - 2*side
		return [2]int{1*side + delta, 2*side - 1}, W
	case next[1] == 2*side && 1*side <= next[0] && next[0] < 2*side:
		delta := next[0] - 1*side
		return [2]int{1*side - 1, 2*side + delta}, N
	case next[1] == 2*side && 2*side <= next[0] && next[0] < 3*side:
		delta := next[0] - 2*side
		return [2]int{1*side - 1 - delta, 3*side - 1}, W
	case next[0] == 2*side-1 && 0*side <= next[1] && next[1] < 1*side:
		delta := next[1] - 0*side
		return [2]int{1*side + delta, 1 * side}, E
	case next[1] == 0*side-1 && 2*side <= next[0] && next[0] < 3*side:
		delta := next[0] - 2*side
		return [2]int{1*side - 1 - delta, 1 * side}, E
	case next[1] == 0*side-1 && 3*side <= next[0] && next[0] < 4*side:
		delta := next[0] - 3*side
		return [2]int{0 * side, 1*side + delta}, S
	case next[0] == 4*side && 0*side <= next[1] && next[1] < 1*side:
		delta := next[1] - 0*side
		return [2]int{0 * side, 2*side + delta}, S
	case next[0] == 3*side && 1*side <= next[1] && next[1] < 2*side:
		delta := next[1] - 1*side
		return [2]int{3*side + delta, 1*side - 1}, W
	case next[1] == 1*side && 3*side <= next[0] && next[0] < 4*side:
		delta := next[0] - 3*side
		return [2]int{3*side - 1, 1*side + delta}, N
	default:
		panic(next)
	}
}

func (p *p) Solve2() any {
	util.Assert(p.lenI == 4*50, "lenI != 4*50")
	util.Assert(p.lenJ == 3*50, "lenJ != 3*50")

	cur := p.startPos
	dir := E

	for pc, nStep := range p.dists {
		for step := int64(0); step < nStep; step++ {
			next := [2]int{cur[0] + deltas[dir][0], cur[1] + deltas[dir][1]}
			nextDir := dir
			if !p.inbound(next[0], next[1]) || p.grids[next[0]][next[1]] == void {
				next, nextDir = p.wrap2(next)
			}
			if p.grids[next[0]][next[1]] == wall {
				break
			}
			cur, dir = next, nextDir
		}
		if pc != len(p.dists)-1 {
			dir = dir.turn(p.turns[pc])
		}
	}

	return 1000*(cur[0]+1) + 4*(cur[1]+1) + int(dir)
}

func init() {
	base.Register(22, &p{})
}

type direction uint8

const (
	E direction = iota
	S
	W
	N
)

var deltas = map[direction][2]int{
	E: {0, 1},
	S: {1, 0},
	W: {0, -1},
	N: {-1, 0},
}

func (d direction) turn(left bool) direction {
	if left {
		return direction((d + 3) % 4)
	} else {
		return direction((d + 1) % 4)
	}
}
