package puzzle14

import (
	"advent2022/base"
	"advent2022/util"
	"fmt"
	"maps"
	"strings"
)

type p struct {
	lines []string
	rocks map[[2]int]bool
	maxY  int
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
	p.rocks = make(map[[2]int]bool)

	scan := func(str string) [2]int {
		var x, y int
		util.Assert(
			util.Must(fmt.Sscanf(str, "%d,%d", &x, &y)) == 2,
			"invalid input: "+str,
		)
		p.maxY = max(p.maxY, y)
		return [2]int{x, y}
	}

	for _, line := range p.lines {
		points := strings.Split(line, " -> ")
		curr := scan(points[0])
		p.rocks[curr] = true
		for _, point := range points[1:] {
			next := scan(point)
			if next[0] == curr[0] {
				for y := min(curr[1], next[1]); y <= max(curr[1], next[1]); y++ {
					p.rocks[[2]int{curr[0], y}] = true
				}
			} else {
				for x := min(curr[0], next[0]); x <= max(curr[0], next[0]); x++ {
					p.rocks[[2]int{x, curr[1]}] = true
				}
			}
			curr = next
		}
	}
}

func (p *p) Solve1() any {
	occupied := maps.Clone(p.rocks)
	count := 0

	for {
		next := [2]int{500, 0}
		count++
		for {
			if next[1] > p.maxY {
				break
			}
			if !occupied[[2]int{next[0], next[1] + 1}] {
				next[1]++
				continue
			}
			if !occupied[[2]int{next[0] - 1, next[1] + 1}] {
				next[0]--
				next[1]++
				continue
			}
			if !occupied[[2]int{next[0] + 1, next[1] + 1}] {
				next[0]++
				next[1]++
				continue
			}
			occupied[next] = true
			break
		}
		if next[1] > p.maxY {
			break
		}
	}

	return count - 1
}

func (p *p) Solve2() any {
	occupied := maps.Clone(p.rocks)

	count := 0

	for {
		next := [2]int{500, 0}
		count++
		for {
			if next[1] == p.maxY+1 {
				occupied[next] = true
				break
			}
			if !occupied[[2]int{next[0], next[1] + 1}] {
				next[1]++
				continue
			}
			if !occupied[[2]int{next[0] - 1, next[1] + 1}] {
				next[0]--
				next[1]++
				continue
			}
			if !occupied[[2]int{next[0] + 1, next[1] + 1}] {
				next[0]++
				next[1]++
				continue
			}
			occupied[next] = true
			break
		}
		if next == [2]int{500, 0} {
			break
		}
	}

	return count
}

func init() {
	base.Register(14, &p{})
}
