package puzzle18

import (
	"advent2022/base"
	"advent2022/util"
	"strings"
)

type p struct {
	cubes      [][3]int64
	minI, maxI int64
	minJ, maxJ int64
	minK, maxK int64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.cubes = make([][3]int64, 0, len(lines))
	for _, line := range lines {
		cord := util.ArrayStrToInt64(strings.SplitN(line, ",", 3))
		p.cubes = append(p.cubes, [3]int64{cord[0], cord[1], cord[2]})
		p.updateMinMax(cord)
	}
}

func (p *p) updateMinMax(cord []int64) {
	p.minI = min(p.minI, cord[0])
	p.maxI = max(p.maxI, cord[0])
	p.minJ = min(p.minJ, cord[1])
	p.maxJ = max(p.maxJ, cord[1])
	p.minK = min(p.minK, cord[2])
	p.maxK = max(p.maxK, cord[2])
}

func (p *p) inbound(cord [3]int64) bool {
	return p.minI-1 <= cord[0] && cord[0] <= p.maxI+1 &&
		p.minJ-1 <= cord[1] && cord[1] <= p.maxJ+1 &&
		p.minK-1 <= cord[2] && cord[2] <= p.maxK+1
}

func (p *p) Solve1() any {
	exist := map[[3]int64]bool{}
	countConnected := func(cord [3]int64) int {
		return util.Reduce(deltas[:], func(acc int, delta [3]int64) int {
			if exist[[3]int64{cord[0] + delta[0], cord[1] + delta[1], cord[2] + delta[2]}] {
				return acc + 1
			}
			return acc
		}, 0)
	}

	connection := 0

	for _, cord := range p.cubes {
		connection += countConnected(cord)
		exist[cord] = true
	}

	return 6*len(p.cubes) - 2*connection
}

func (p *p) Solve2() any {
	exist := util.ToVis(p.cubes)
	vis := map[[3]int64]bool{}
	queue := [][3]int64{{p.minI - 1, p.minJ - 1, p.minK - 1}}

	surface := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if exist[curr] {
			surface++
			continue
		}
		if !vis[curr] {
			vis[curr] = true
			for _, delta := range deltas {
				next := [3]int64{curr[0] + delta[0], curr[1] + delta[1], curr[2] + delta[2]}
				if p.inbound(next) {
					queue = append(queue, next)
				}
			}
		}
	}

	return surface
}

func init() {
	base.Register(18, &p{})
}

var deltas = [6][3]int64{
	{0, 0, 1},
	{0, 0, -1},
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
}
