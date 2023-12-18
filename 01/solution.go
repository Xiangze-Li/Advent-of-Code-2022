package puzzle01

import (
	"advent2022/base"
	"advent2022/util"
	"slices"
)

type p struct {
	blocks [][]string
	values []int64
}

func (p *p) Init(filename string) {
	p.blocks = util.GetBlocks(filename)
	p.values = make([]int64, 0, len(p.blocks))
	for i := 0; i < len(p.blocks); i++ {
		var sum int64
		for _, v := range util.ArrayStrToInt64(p.blocks[i]) {
			sum += v
		}
		p.values = append(p.values, sum)
	}
}

func (p *p) Solve1() any {
	return slices.Max(p.values)
}

func (p *p) Solve2() any {
	slices.Sort(p.values)
	return p.values[len(p.values)-1] + p.values[len(p.values)-2] + p.values[len(p.values)-3]
}

func init() {
	base.Register(1, &p{})
}
