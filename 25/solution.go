package puzzle25

import (
	"advent2022/base"
	"advent2022/util"
)

type p struct {
	nums []int64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.nums = make([]int64, 0, len(lines))

	for _, line := range lines {
		p.nums = append(p.nums, util.Must(util.FromBalancedQuinary(line)))
	}
}

func (p *p) Solve1() any {
	sum := util.Reduce(p.nums, func(sum, next int64) int64 { return sum + next }, 0)

	return util.ToBalancedQuinary(sum)
}

func (p *p) Solve2() any {
	return 0
}

func init() {
	base.Register(25, &p{})
}
