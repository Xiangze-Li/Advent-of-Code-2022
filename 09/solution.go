package puzzle09

import (
	"advent2022/base"
	"advent2022/util"
	"strconv"
)

type p struct {
	dirs  []byte
	dists []int
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.dirs = make([]byte, 0, len(lines))
	p.dists = make([]int, 0, len(lines))
	for _, line := range lines {
		p.dirs = append(p.dirs, line[0])
		p.dists = append(p.dists, util.Must(strconv.Atoi(line[2:])))
	}
}

func tailFollows(headX, headY, tailX, tailY int) (newX, newY int) {
	newX, newY = tailX, tailY

	diffX, diffY := headX-tailX, headY-tailY

	if util.Abs(diffX) <= 1 && util.Abs(diffY) <= 1 {
		return
	}

	newX += util.Sign(diffX)
	newY += util.Sign(diffY)
	return
}

func (p *p) Solve1() any {
	vis := map[[2]int]bool{}

	headX, headY := 0, 0
	tailX, tailY := 0, 0
	vis[[2]int{tailX, tailY}] = true

	for i, dir := range p.dirs {
		delta := deltas[dir]
		dist := p.dists[i]

		for j := 0; j < dist; j++ {
			headX, headY = headX+delta[0], headY+delta[1]
			tailX, tailY = tailFollows(headX, headY, tailX, tailY)
			vis[[2]int{tailX, tailY}] = true
		}
	}

	return util.ReduceMap(vis, func(acc int64, _ [2]int, v bool) int64 {
		if v {
			return acc + 1
		}
		return acc
	}, 0)
}

func (p *p) Solve2() any {
	vis := map[[2]int]bool{}

	rope := [10][2]int{}
	vis[rope[9]] = true

	for i, dir := range p.dirs {
		delta := deltas[dir]
		dist := p.dists[i]

		for j := 0; j < dist; j++ {
			rope[0][0] += delta[0]
			rope[0][1] += delta[1]
			for k := 1; k < 10; k++ {
				rope[k][0], rope[k][1] = tailFollows(rope[k-1][0], rope[k-1][1], rope[k][0], rope[k][1])
			}
			vis[rope[9]] = true
		}
	}

	return util.ReduceMap(vis, func(acc int64, _ [2]int, v bool) int64 {
		if v {
			return acc + 1
		}
		return acc
	}, 0)
}

func init() {
	base.Register(9, &p{})
}

var deltas = map[byte][2]int{
	'U': {-1, 0},
	'D': {1, 0},
	'L': {0, -1},
	'R': {0, 1},
}
