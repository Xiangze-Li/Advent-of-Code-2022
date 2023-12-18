package puzzle21

import (
	"advent2022/base"
	"advent2022/util"
	"maps"
	"strconv"
	"strings"
)

type p struct {
	values map[string]int64
	exprs  map[string]expression
}

type expression struct {
	to         string
	oprL, oprR string
	op         byte
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)

	p.values = make(map[string]int64, len(lines))
	p.exprs = make(map[string]expression, len(lines))

	for _, line := range lines {
		sp := strings.SplitN(line, ": ", 2)
		if val, err := strconv.ParseInt(sp[1], 10, 64); err == nil {
			p.values[sp[0]] = val
		} else {
			spp := strings.SplitN(sp[1], " ", 3)
			p.exprs[sp[0]] = expression{sp[0], spp[0], spp[2], spp[1][0]}
		}
	}
}

func (p *p) eval1(pValues *map[string]int64, target string) (ok bool, val int64, pending []expression) {
	if val, ok = (*pValues)[target]; ok {
		return
	}

	okL, valL, pendingL := p.eval1(pValues, p.exprs[target].oprL)
	okR, valR, pendingR := p.eval1(pValues, p.exprs[target].oprR)
	if okL && okR {
		switch p.exprs[target].op {
		case '+':
			val = valL + valR
		case '*':
			val = valL * valR
		case '-':
			val = valL - valR
		case '/':
			val = valL / valR
		}
		(*pValues)[target] = val
		return true, val, nil
	}
	if okL {
		return false, 0, append([]expression{p.exprs[target]}, pendingR...)
	}
	if okR {
		return false, 0, append([]expression{p.exprs[target]}, pendingL...)
	}
	panic("both not ok")
}

func (p *p) Solve1() any {
	values := maps.Clone(p.values)

	_, val, _ := p.eval1(&values, "root")
	return val
}

func (p *p) eval2(pValues *map[string]int64, target string) (ok bool, val int64, pending []expression) {
	if target == "humn" {
		return false, 0, []expression{{to: "humn"}}
	}

	if val, ok = (*pValues)[target]; ok {
		return
	}

	okL, valL, pendingL := p.eval2(pValues, p.exprs[target].oprL)
	okR, valR, pendingR := p.eval2(pValues, p.exprs[target].oprR)
	if okL && okR {
		switch p.exprs[target].op {
		case '+':
			val = valL + valR
		case '*':
			val = valL * valR
		case '-':
			val = valL - valR
		case '/':
			val = valL / valR
		}
		(*pValues)[target] = val
		return true, val, nil
	}
	if okL {
		return false, 0, append([]expression{p.exprs[target]}, pendingR...)
	}
	if okR {
		return false, 0, append([]expression{p.exprs[target]}, pendingL...)
	}
	panic("both not ok")
}

func (p *p) Solve2() any {
	values := maps.Clone(p.values)
	delete(values, "humn")

	_, _, pending := p.eval2(&values, "root")

	for _, expr := range pending {
		if expr.to == "root" {
			if val, okL := values[expr.oprL]; okL {
				values[expr.oprR] = val
			} else {
				values[expr.oprL] = values[expr.oprR]
			}
			continue
		}
		if expr.to == "humn" {
			break
		}

		if valL, okL := values[expr.oprL]; okL {
			switch expr.op {
			case '+':
				values[expr.oprR] = values[expr.to] - valL
			case '-':
				values[expr.oprR] = valL - values[expr.to]
			case '*':
				values[expr.oprR] = values[expr.to] / valL
			case '/':
				values[expr.oprR] = valL / values[expr.to]
			}
		} else {
			valR, ok := values[expr.oprR]
			util.Assert(ok, "both not ok")
			switch expr.op {
			case '+':
				values[expr.oprL] = values[expr.to] - valR
			case '-':
				values[expr.oprL] = values[expr.to] + valR
			case '*':
				values[expr.oprL] = values[expr.to] / valR
			case '/':
				values[expr.oprL] = values[expr.to] * valR
			}
		}
	}

	return values["humn"]
}

func init() {
	base.Register(21, &p{})
}

func (p *p) topoSort() []string {
	order := make([]string, 0, len(p.exprs))
	indeg := make(map[string]int, len(p.exprs))

	toTable := map[string][]string{}

	for k, expr := range p.exprs {
		in := 2
		toTable[expr.oprL] = append(toTable[expr.oprL], k)
		toTable[expr.oprR] = append(toTable[expr.oprR], k)
		if _, ok := p.values[expr.oprL]; ok {
			in--
		}
		if _, ok := p.values[expr.oprR]; ok {
			in--
		}
		indeg[k] = in
	}

	for len(indeg) > 0 {
		var k string
		for k = range indeg {
			if indeg[k] == 0 {
				break
			}
		}
		order = append(order, k)
		delete(indeg, k)

		for _, to := range toTable[k] {
			if indeg[to] > 0 {
				indeg[to]--
			}
		}
	}
	return order
}
