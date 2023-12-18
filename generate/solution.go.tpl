package puzzle{{.DayString}}

import (
	"advent2022/base"
	"advent2022/util"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

func (p *p) Solve1() int64 {
	return 0
}

func (p *p) Solve2() int64 {
	return 0
}

func init() {
	base.Register({{.Day}}, &p{})
}
