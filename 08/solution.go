package puzzle08

import (
	"advent2022/base"
	"advent2022/util"
)

type p struct {
	grid [][]int64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.grid = util.SliceND[int64](len(lines), len(lines[0])).([][]int64)
	for i, line := range lines {
		for j, c := range line {
			p.grid[i][j] = int64(c - '0')
		}
	}
}

func (p *p) Solve1() any {
	lenI, lenJ := len(p.grid), len(p.grid[0])
	visible := util.SliceND[bool](lenI, lenJ).([][]bool)

	for i := 0; i < lenI; i++ {
		var maxL, maxR int64 = -1, -1
		for j := 0; j < lenJ; j++ {
			jj := lenJ - 1 - j
			if p.grid[i][j] > maxL {
				maxL = p.grid[i][j]
				visible[i][j] = true
			}
			if p.grid[i][jj] > maxR {
				maxR = p.grid[i][jj]
				visible[i][jj] = true
			}
		}
	}

	for j := 0; j < lenJ; j++ {
		var maxU, maxD int64 = -1, -1
		for i := 0; i < lenI; i++ {
			ii := lenI - 1 - i
			if p.grid[i][j] > maxU {
				maxU = p.grid[i][j]
				visible[i][j] = true
			}
			if p.grid[ii][j] > maxD {
				maxD = p.grid[ii][j]
				visible[ii][j] = true
			}
		}
	}

	return util.Reduce(visible, func(acc int64, row []bool) int64 {
		return util.Reduce(row, func(acc int64, b bool) int64 {
			if b {
				return acc + 1
			}
			return acc
		}, acc)
	}, 0)
}

func (p *p) Solve2() any {
	maxium := 0
	lenI, lenJ := len(p.grid), len(p.grid[0])

	for i := 1; i < lenI-1; i++ {
		for j := 1; j < lenJ-1; j++ {
			if p.grid[i][j] == 0 {
				continue
			}
			score := 1
			for _, delta := range deltas {
				ii, jj := i, j
				dist := 0
				for {
					ii += delta[0]
					jj += delta[1]
					if ii < 0 || ii >= lenI || jj < 0 || jj >= lenJ {
						break
					}
					dist++
					if p.grid[ii][jj] >= p.grid[i][j] {
						break
					}
				}
				score *= dist
				if dist == 0 {
					break
				}
			}
			maxium = max(maxium, score)
		}
	}

	return maxium
}

func init() {
	base.Register(8, &p{})
}

var deltas = [...][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}
