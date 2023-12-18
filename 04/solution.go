package puzzle04

import (
	"advent2022/base"
	"advent2022/util"
	"regexp"
	"strconv"
)

type range_ struct {
	lb, ub int
}

type p struct {
	lines []string
	pairs [][2]range_
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
	p.pairs = make([][2]range_, 0, len(p.lines))
	re := regexp.MustCompile(`[,\-]`)
	for _, line := range p.lines {
		sp := re.Split(line, -1)
		util.Assert(len(sp) == 4, "bad line: "+line)

		p.pairs = append(p.pairs, [2]range_{
			{util.Must(strconv.Atoi(sp[0])), util.Must(strconv.Atoi(sp[1]))},
			{util.Must(strconv.Atoi(sp[2])), util.Must(strconv.Atoi(sp[3]))},
		})
	}
}

func (p *p) Solve1() any {
	return util.Reduce(p.pairs, func(acc int64, pair [2]range_) int64 {
		if (pair[0].lb <= pair[1].lb && pair[1].ub <= pair[0].ub) ||
			(pair[1].lb <= pair[0].lb && pair[0].ub <= pair[1].ub) {
			return acc + 1
		}
		return acc
	}, 0)
}

func (p *p) Solve2() any {
	return util.Reduce(p.pairs, func(acc int64, pair [2]range_) int64 {
		if pair[0].ub < pair[1].lb || pair[1].ub < pair[0].lb {
			return acc - 1
		}
		return acc
	}, int64(len(p.pairs)))
}

func init() {
	base.Register(4, &p{})
}
