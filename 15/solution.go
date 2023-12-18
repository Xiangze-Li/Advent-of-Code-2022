package puzzle15

import (
	"advent2022/base"
	"advent2022/util"
	"cmp"
	"fmt"
	"slices"
)

type p struct {
	lines   []string
	banned  map[int][][2]int
	beacons map[[2]int]bool

	targetY int
	lb, ub  int
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
	p.banned = make(map[int][][2]int)
	p.beacons = make(map[[2]int]bool)

	for _, line := range p.lines {
		var senX, senY, bcnX, bcnY int
		util.Must(fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &senX, &senY, &bcnX, &bcnY))
		p.beacons[[2]int{bcnX, bcnY}] = true
		dist := util.Abs(bcnX-senX) + util.Abs(bcnY-senY)
		for deltaY := -dist; deltaY <= dist; deltaY++ {
			deltaX := dist - util.Abs(deltaY)
			lb := senX - deltaX
			ub := senX + deltaX
			p.banned[senY+deltaY] = append(p.banned[senY+deltaY], [2]int{lb, ub + 1})
		}
	}

	p.targetY = 2_000_000
	p.lb = 0
	p.ub = 4_000_001

	if filename == "15/example.txt" {
		p.targetY = 10
		p.lb = 0
		p.ub = 21
	}
}

func (p *p) Solve1() any {

	line := p.banned[p.targetY]
	merged := make([][2]int, 0, len(line))

	slices.SortFunc(line, func(a, b [2]int) int {
		if c := cmp.Compare(a[0], b[0]); c != 0 {
			return c
		}
		return cmp.Compare(a[1], b[1])
	})
	for _, v := range line {
		if len(merged) == 0 || merged[len(merged)-1][1] < v[0] {
			merged = append(merged, v)
		} else if merged[len(merged)-1][1] < v[1] {
			merged[len(merged)-1][1] = v[1]
		}
	}

	return util.Reduce(merged, func(acc int, v [2]int) int {
		return acc + v[1] - v[0]
	}, util.ReduceMap(p.beacons, func(acc int, bcn [2]int, _ bool) int {
		if bcn[1] == p.targetY {
			return acc - 1
		}
		return acc
	}, 0))
}

func (p *p) Solve2() any {
loopY:
	for y := p.lb; y < p.ub; y++ {
		line := p.banned[y]
		merged := make([][2]int, 0, len(line))

		slices.SortFunc(line, func(a, b [2]int) int {
			if c := cmp.Compare(a[0], b[0]); c != 0 {
				return c
			}
			return cmp.Compare(a[1], b[1])
		})
		for _, v := range line {
			if len(merged) == 0 || merged[len(merged)-1][1] < v[0] {
				merged = append(merged, v)
			} else if merged[len(merged)-1][1] < v[1] {
				merged[len(merged)-1][1] = v[1]
			}
		}

		for _, v := range merged {
			if v[1] < p.lb && p.ub < v[0] {
				continue
			}
			if v[0] <= p.lb && p.ub <= v[1] {
				continue loopY
			}
			if v[0] <= p.lb {
				return v[1]*(p.ub-1) + y
			}
			return (v[0]-1)*(p.ub-1) + y
		}
	}

	return -1
}

func init() {
	base.Register(15, &p{})
}
