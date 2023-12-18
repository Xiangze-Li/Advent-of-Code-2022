package puzzle12

import (
	"advent2022/base"
	"advent2022/util"
	"strings"
)

type p struct {
	grid  [][]byte
	start [2]int
	end   [2]int
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.grid = make([][]byte, len(lines))
	for i, line := range lines {
		p.grid[i] = []byte(lines[i])
		if idx := strings.Index(line, "S"); idx != -1 {
			p.start = [2]int{i, idx}
			p.grid[i][idx] = 'a'
		}
		if idx := strings.Index(line, "E"); idx != -1 {
			p.end = [2]int{i, idx}
			p.grid[i][idx] = 'z'
		}
	}
}

func (p *p) Solve1() any {
	lenI, lenJ := len(p.grid), len(p.grid[0])
	inbound := func(i, j int) bool {
		return i >= 0 && i < lenI && j >= 0 && j < lenJ
	}
	dist := util.SliceND[int](lenI, lenJ).([][]int)

	dist[p.start[0]][p.start[1]] = 1
	queue := [][2]int{p.start}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, dir := range [...][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			ni, nj := i+dir[0], j+dir[1]
			if !inbound(ni, nj) || dist[ni][nj] != 0 ||
				(p.grid[ni][nj] > p.grid[i][j] && p.grid[ni][nj]-p.grid[i][j] > 1) {
				continue
			}
			dist[ni][nj] = dist[i][j] + 1
			if ni == p.end[0] && nj == p.end[1] {
				return dist[i][j]
			}
			queue = append(queue, [2]int{ni, nj})
		}
	}

	return -1
}

func (p *p) Solve2() any {
	lenI, lenJ := len(p.grid), len(p.grid[0])
	inbound := func(i, j int) bool {
		return i >= 0 && i < lenI && j >= 0 && j < lenJ
	}
	dist := util.SliceND[int](lenI, lenJ).([][]int)

	dist[p.end[0]][p.end[1]] = 1
	queue := [][2]int{p.end}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, dir := range [...][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			ni, nj := i+dir[0], j+dir[1]
			if !inbound(ni, nj) || dist[ni][nj] != 0 ||
				(p.grid[ni][nj] < p.grid[i][j] && p.grid[i][j]-p.grid[ni][nj] > 1) {
				continue
			}
			dist[ni][nj] = dist[i][j] + 1
			if p.grid[ni][nj] == 'a' {
				return dist[i][j]
			}
			queue = append(queue, [2]int{ni, nj})
		}
	}

	return -1
}

func init() {
	base.Register(12, &p{})
}
