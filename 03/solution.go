package puzzle03

import (
	"advent2022/base"
	"advent2022/util"
	"unicode"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

func (p *p) Solve1() any {
	return util.Reduce(p.lines, func(sum int64, line string) int64 {
		length := len(line)
		partA, partB := line[:length/2], line[length/2:]

		existA := util.Associative([]byte(partA), func(c byte) (byte, bool) {
			return c, true
		})

		for _, c := range []byte(partB) {
			if existA[c] {
				if unicode.IsLower(rune(c)) {
					return sum + int64(c-'a'+1)
				} else {
					return sum + int64(c-'A'+27)
				}
			}
		}

		return sum
	}, 0)
}

func (p *p) Solve2() any {
	var sum int64

	for i := 0; i+2 < len(p.lines); i += 3 {
		existA := util.Associative([]byte(p.lines[i]), func(c byte) (byte, bool) {
			return c, true
		})
		existB := util.Associative([]byte(p.lines[i+1]), func(c byte) (byte, bool) {
			return c, true
		})
		existC := util.Associative([]byte(p.lines[i+2]), func(c byte) (byte, bool) {
			return c, true
		})

		for score := int64(0); score < 26*2; score++ {
			c := byte(score) + 'a'
			if score >= 26 {
				c = byte(score-26) + 'A'
			}
			if existA[c] && existB[c] && existC[c] {
				sum += score + 1
			}
		}

	}

	return sum
}

func init() {
	base.Register(3, &p{})
}
