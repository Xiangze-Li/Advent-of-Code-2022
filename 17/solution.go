package puzzle17

import (
	"advent2022/base"
	"advent2022/util"
)

const T, F bool = true, false

type p struct {
	ops []byte
}

var rocks = [5][]byte{
	{
		0b00001111,
	},
	{
		0b00000010,
		0b00000111,
		0b00000010,
	},
	{
		0b00000111,
		0b00000100,
		0b00000100,
	},
	{
		0b00000001,
		0b00000001,
		0b00000001,
		0b00000001,
	},
	{
		0b00000011,
		0b00000011,
	},
}

var rockWs = [5]int{4, 3, 3, 1, 2}

func (p *p) Init(filename string) {
	p.ops = []byte(util.GetLines(filename)[0])
}

func (p *p) sim(round int) int {
	opCount := 0
	occupied := []byte{}
	intersects := func(i, j int, rock []byte) bool {
		for ii, row := range rock {
			if i+ii >= len(occupied) {
				return false
			}
			if (row<<j)&occupied[i+ii] != 0 {
				return true
			}
		}
		return false
	}

	calcDepth := func() [7]int {
		res := [7]int{}
		for i := 0; i < 7; i++ {
			m := byte(1 << i)
			for j := len(occupied) - 1; j >= 0; j-- {
				if occupied[j]&m != 0 {
					res[i] = len(occupied) - 1 - j
					break
				}
			}
		}
		return res
	}

	cycleFound := false
	vis := map[[9]int][2]int{}
	extraHeight := 0

	for r := 0; r < round; r++ {
		rock := rocks[r%5]
		rockH, rockW := len(rock), rockWs[r%5]
		i, j := len(occupied)+3, 2
		for {
			op := p.ops[opCount%len(p.ops)]
			opCount++
			if op == '<' {
				if j > 0 && !intersects(i, j-1, rock) {
					j--
				}
			} else {
				if j+rockW < 7 && !intersects(i, j+1, rock) {
					j++
				}
			}

			if i > 0 && !intersects(i-1, j, rock) {
				i--
				continue
			}

			rockTop := i + rockH
			if top := len(occupied); rockTop > top {
				occupied = append(occupied, make([]byte, rockTop-top)...)
			}
			for ii, row := range rock {
				occupied[i+ii] |= row << j
			}
			break
		}

		if !cycleFound {
			key := [9]int{r % 5, opCount % len(p.ops)}
			{
				depthMap := calcDepth()
				copy(key[2:9], depthMap[:])
			}
			if state, ok := vis[key]; ok {
				cycleFound = true
				period := r - state[0]
				height := len(occupied) - state[1]

				fullCycle := (round - r - 1) / period
				r += fullCycle * period
				extraHeight = fullCycle * height
			} else {
				vis[key] = [2]int{r, len(occupied)}
			}
		}
	}

	return len(occupied) + extraHeight
}

func (p *p) Solve1() any {
	return p.sim(2022)
}

func (p *p) Solve2() any {
	return p.sim(1_000_000_000_000)
}

func init() {
	base.Register(17, &p{})
}
