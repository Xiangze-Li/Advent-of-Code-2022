package puzzle06

import (
	"advent2022/base"
	"advent2022/util"
)

type p struct {
	chars []byte
}

func (p *p) Init(filename string) {
	p.chars = []byte(util.GetLines(filename)[0])
}

func (p *p) solve(length int) any {
	m := map[byte]int{}
	for i := 0; i < length; i++ {
		m[p.chars[i]]++
	}
	for i := length; i < len(p.chars); i++ {
		noDup := true
		for _, v := range m {
			if v > 1 {
				noDup = false
				break
			}
		}
		if noDup {
			return i
		}

		m[p.chars[i]]++
		m[p.chars[i-length]]--
		if m[p.chars[i-length]] == 0 {
			delete(m, p.chars[i-length])
		}
	}
	return -1
}

func (p *p) Solve1() any {
	return p.solve(4)
}

func (p *p) Solve2() any {
	return p.solve(14)
}

func init() {
	base.Register(6, &p{})
}
