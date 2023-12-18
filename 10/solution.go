package puzzle10

import (
	"advent2022/base"
	"advent2022/util"
	"strconv"
	"strings"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

func (p *p) Solve1() any {
	var x, cycle int64 = 1, 1
	var sum int64

	for _, line := range p.lines {
		sp := strings.SplitN(line, " ", 2)
		switch sp[0] {
		case "noop": // noop
			if cycle%40 == 20 {
				sum += x * cycle
			}
			cycle++
		case "addx":
			d := util.Must(strconv.ParseInt(sp[1], 10, 64))
			if cycle%40 == 19 {
				sum += x * (cycle + 1)
			} else if cycle%40 == 20 {
				sum += x * cycle
			}
			cycle += 2
			x += d
		}
	}

	return sum
}

func (p *p) Solve2() any {
	var x, cycle int = 1, 0

	screen := [6][40]byte{}
	handleCycle := func(x, cycle int) {
		line := cycle / 40
		col := cycle % 40

		if util.Abs(x-col) <= 1 {
			screen[line][col] = '#'
		} else {
			screen[line][col] = ' '
		}
	}

	for _, line := range p.lines {
		sp := strings.SplitN(line, " ", 2)
		switch sp[0] {
		case "noop": // noop
			handleCycle(x, cycle)
			cycle++
		case "addx":
			d := util.Must(strconv.Atoi(sp[1]))
			handleCycle(x, cycle)
			handleCycle(x, cycle+1)
			cycle += 2
			x += d
		}
	}

	res := [6]string{}
	for i := 0; i < 6; i++ {
		res[i] = string(screen[i][:])
	}
	return res
}

func init() {
	base.Register(10, &p{})
}
