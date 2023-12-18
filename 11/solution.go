package puzzle11

import (
	"advent2022/base"
	"advent2022/util"
	"slices"
	"strconv"
	"strings"
)

type p struct {
	monkeys []*monkey
}

type monkey struct {
	items       []int64
	op          func(int64) int64
	test        int64
	trueBranch  int
	falseBranch int
}

func generateOp(expr []string) func(int64) int64 {
	util.Assert(len(expr) == 3, "invalid expr")
	return func(x int64) int64 {
		oprL, oprR := x, x
		if expr[0] != "old" {
			oprL = util.Must(strconv.ParseInt(expr[0], 10, 64))
		}
		if expr[2] != "old" {
			oprR = util.Must(strconv.ParseInt(expr[2], 10, 64))
		}
		switch expr[1] {
		case "+":
			return oprL + oprR
		case "-":
			return oprL - oprR
		case "*":
			return oprL * oprR
		case "/":
			return oprL / oprR
		default:
			util.Assert(false, "invalid op")
			return 0
		}
	}

}

func (p *p) Init(filename string) {
	blocks := util.GetBlocks(filename)
	p.monkeys = make([]*monkey, 0, len(blocks))

	for _, block := range blocks {
		util.Assert(len(block) == 6, "invalid block")

		m := &monkey{}

		m.items = util.ArrayStrToInt64(
			strings.Split(
				strings.Split(block[1], ": ")[1],
				", ",
			))
		{
			opStr := strings.Split(block[2], "=")[1]
			expr := strings.Fields(opStr)
			m.op = generateOp(expr)
		}
		m.test = util.Must(strconv.ParseInt(strings.Split(block[3], "by ")[1], 10, 64))
		{
			f := strings.Fields(block[4])
			m.trueBranch = util.Must(strconv.Atoi(f[len(f)-1]))
		}
		{
			f := strings.Fields(block[5])
			m.falseBranch = util.Must(strconv.Atoi(f[len(f)-1]))
		}

		p.monkeys = append(p.monkeys, m)
	}
}

func (p *p) Solve1() any {
	count := make([]int64, len(p.monkeys))
	monkeys := make([]*monkey, len(p.monkeys))
	for i, m := range p.monkeys {
		mm := *m
		mm.items = slices.Clone(m.items)
		monkeys[i] = &(mm)
	}

	for round := 0; round < 20; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				count[i]++
				newItem := m.op(item) / 3
				if newItem%m.test == 0 {
					monkeys[m.trueBranch].items = append(monkeys[m.trueBranch].items, newItem)
				} else {
					monkeys[m.falseBranch].items = append(monkeys[m.falseBranch].items, newItem)
				}
			}
			m.items = []int64{}
		}
	}

	slices.Sort(count)
	return count[len(count)-1] * count[len(count)-2]
}

func (p *p) Solve2() any {
	count := make([]int64, len(p.monkeys))
	monkeys := make([]*monkey, len(p.monkeys))

	var modd int64 = 1
	for i, m := range p.monkeys {
		mm := *m
		mm.items = slices.Clone(m.items)
		monkeys[i] = &(mm)
		modd = util.LCM(modd, m.test)
	}

	for round := 0; round < 10000; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				count[i]++
				newItem := m.op(item) % modd
				if newItem%m.test == 0 {
					monkeys[m.trueBranch].items = append(monkeys[m.trueBranch].items, newItem)
				} else {
					monkeys[m.falseBranch].items = append(monkeys[m.falseBranch].items, newItem)
				}
			}
			m.items = []int64{}
		}
	}

	slices.Sort(count)
	return count[len(count)-1] * count[len(count)-2]
}

func init() {
	base.Register(11, &p{})
}
