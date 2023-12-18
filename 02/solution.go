package puzzle02

import (
	"advent2022/base"
	"advent2022/util"
	"fmt"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

func (p *p) Solve1() any {
	score := func(op, my rune) int64 {
		var s int64
		s = int64(my-'X') + 1
		switch (my - 'X') - (op - 'A') {
		case 0:
			s += 3
		case 1, -2:
			s += 6
		case -1, 2:
			s += 0
		}
		return s
	}

	return util.Reduce(p.lines, func(sum int64, line string) int64 {
		var op, my rune
		fmt.Sscanf(line, "%c %c", &op, &my)
		return sum + score(op, my)
	}, 0)
}

func (p *p) Solve2() any {
	score := func(op, outcome rune) int64 {
		outcomeInt := int64(outcome - 'X')
		opInt := int64(op - 'A')
		var s int64 = outcomeInt * 3
		switch outcomeInt {
		case 0:
			s += (opInt-1+3)%3 + 1
		case 1:
			s += opInt + 1
		case 2:
			s += (opInt+1)%3 + 1
		}
		return s
	}

	return util.Reduce(p.lines, func(sum int64, line string) int64 {
		var op, outcome rune
		fmt.Sscanf(line, "%c %c", &op, &outcome)
		return sum + score(op, outcome)
	}, 0)
}

func init() {
	base.Register(2, &p{})
}
