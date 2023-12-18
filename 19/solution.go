package puzzle19

import (
	"advent2022/base"
	"advent2022/util"
	"regexp"
)

var re = regexp.MustCompile(`\d+`)

type p struct {
	plans []plan
}

type plan struct {
	no            int64
	oreCostOre    int64
	clayCostOre   int64
	obsdCostOre   int64
	obsdCostClay  int64
	geodeCostOre  int64
	geodeCostObsd int64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.plans = make([]plan, 0, len(lines))
	for _, line := range lines {
		match := re.FindAllString(line, -1)
		util.Assert(len(match) == 7, "invalid line: "+line)
		numbers := util.ArrayStrToInt64(match)
		p.plans = append(p.plans, plan{
			no:            numbers[0],
			oreCostOre:    numbers[1],
			clayCostOre:   numbers[2],
			obsdCostOre:   numbers[3],
			obsdCostClay:  numbers[4],
			geodeCostOre:  numbers[5],
			geodeCostObsd: numbers[6],
		})
	}
}

type state struct {
	time                               int64
	ore, clay, obsd, geode             int64
	oreBot, clayBot, obsdBot, geodeBot int64
}

var globalBest int64

func dfs(p plan, cur state) int64 {
	if cur.time == 0 {
		return 0
	}
	if globalBest > cur.geode+(2*cur.geodeBot+cur.time-1)*cur.time/2 {
		return 0
	}
	if cur.oreBot >= p.geodeCostOre && cur.obsdBot >= p.geodeCostObsd {
		return (2*cur.geodeBot + cur.time - 1) * cur.time / 2
	}

	oreLimitHit := cur.oreBot >= max(p.clayCostOre, p.obsdCostOre, p.geodeCostOre)
	clayLimitHit := cur.clayBot >= p.obsdCostClay
	obsdLimitHit := cur.obsdBot >= p.geodeCostObsd

	best := int64(0)
	next := cur
	next.time--
	next.ore += cur.oreBot
	next.clay += cur.clayBot
	next.obsd += cur.obsdBot
	next.geode += cur.geodeBot
	if !oreLimitHit {
		best = max(best, cur.geodeBot+dfs(p, next))
	}
	if cur.ore >= p.oreCostOre && !oreLimitHit {
		nextt := next
		nextt.ore -= p.oreCostOre
		nextt.oreBot++
		best = max(best, cur.geodeBot+dfs(p, nextt))
	}
	if cur.ore >= p.clayCostOre && !clayLimitHit {
		nextt := next
		nextt.ore -= p.clayCostOre
		nextt.clayBot++
		best = max(best, cur.geodeBot+dfs(p, nextt))
	}
	if cur.ore >= p.obsdCostOre && cur.clay >= p.obsdCostClay && !obsdLimitHit {
		nextt := next
		nextt.ore -= p.obsdCostOre
		nextt.clay -= p.obsdCostClay
		nextt.obsdBot++
		best = max(best, cur.geodeBot+dfs(p, nextt))
	}
	if cur.ore >= p.geodeCostOre && cur.obsd >= p.geodeCostObsd {
		nextt := next
		nextt.ore -= p.geodeCostOre
		nextt.obsd -= p.geodeCostObsd
		nextt.geodeBot++
		best = max(best, cur.geodeBot+dfs(p, nextt))
	}

	globalBest = max(globalBest, best)
	return best
}

func (p *p) Solve1() any {
	score := make([]int64, 0, len(p.plans))

	for _, plan := range p.plans {
		globalBest = 0
		score = append(score, dfs(plan, state{time: 24, oreBot: 1}))
	}

	return util.ReduceIndex(score, func(acc int64, i int, cur int64) int64 {
		return acc + int64(i+1)*cur
	}, 0)
}

func (p *p) Solve2() any {
	score := make([]int64, 0, 3)
	for _, plan := range p.plans[:3] {
		globalBest = 0
		score = append(score, dfs(plan, state{time: 32, oreBot: 1}))
	}

	return util.Reduce(score, func(acc int64, cur int64) int64 {
		return acc * cur
	}, 1)
}

func init() {
	base.Register(19, &p{})
}
