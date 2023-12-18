package puzzle24

import (
	"advent2022/base"
	"advent2022/util"
	"slices"
)

type p struct {
	blizzards   []blizzard
	entry, exit [2]int
	lenI, lenJ  int
	cycle       int
}

type blizzard struct {
	loc [2]int
	dir direction
}

func (p *p) next(b blizzard) blizzard {
	next := [2]int{b.loc[0] + deltas[b.dir][0], b.loc[1] + deltas[b.dir][1]}
	if next[0] == 0 {
		next[0] = p.lenI - 2
	}
	if next[0] == p.lenI-1 {
		next[0] = 1
	}
	if next[1] == 0 {
		next[1] = p.lenJ - 2
	}
	if next[1] == p.lenJ-1 {
		next[1] = 1
	}
	return blizzard{next, b.dir}
}

func (p *p) nextAll(bs []blizzard) []blizzard {
	next := make([]blizzard, 0, len(bs))
	for _, b := range bs {
		next = append(next, p.next(b))
	}
	return next
}

func (p *p) inbound(loc [2]int) bool {
	return 0 < loc[0] && loc[0] < p.lenI-1 && 0 < loc[1] && loc[1] < p.lenJ-1
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)

	p.lenI = len(lines)
	p.lenJ = len(lines[0])
	p.cycle = util.LCM(p.lenI-2, p.lenJ-2)

	for i, line := range lines {
		if i == 0 {
			for j, c := range line {
				if c == '.' {
					p.entry = [2]int{i, j}
					break
				}
			}
			continue
		} else if i == p.lenI-1 {
			for j, c := range line {
				if c == '.' {
					p.exit = [2]int{i, j}
					break
				}
			}
			continue
		}

		for j, c := range line {
			switch c {
			case '^':
				p.blizzards = append(p.blizzards, blizzard{[2]int{i, j}, N})
			case 'v':
				p.blizzards = append(p.blizzards, blizzard{[2]int{i, j}, S})
			case '<':
				p.blizzards = append(p.blizzards, blizzard{[2]int{i, j}, W})
			case '>':
				p.blizzards = append(p.blizzards, blizzard{[2]int{i, j}, E})
			}
		}
	}
}

func (p *p) Solve1() any {
	blizs := slices.Clone(p.blizzards)

	t := 0
	possible := map[[2]int]bool{p.entry: true}
	for {
		blizs = p.nextAll(blizs)
		nogo := util.Associative(blizs, func(b blizzard) ([2]int, bool) { return b.loc, true })
		next := make(map[[2]int]bool)
		for loc := range possible {
			for _, delta := range deltas {
				nextLoc := [2]int{loc[0] + delta[0], loc[1] + delta[1]}
				if nextLoc == p.exit {
					return t + 1
				}
				if (nextLoc == p.entry) || (p.inbound(nextLoc) && !nogo[nextLoc]) {
					next[nextLoc] = true
				}
			}
			if !nogo[loc] {
				next[loc] = true
			}
		}
		possible = next
		t++
	}
}

func (p *p) Solve2() any {
	blizs := slices.Clone(p.blizzards)

	t := 0
	possible := map[[2]int]bool{p.entry: true}
	phase := 0
	for {
		blizs = p.nextAll(blizs)
		nogo := util.Associative(blizs, func(b blizzard) ([2]int, bool) { return b.loc, true })
		next := make(map[[2]int]bool)
		for loc := range possible {
			for _, delta := range deltas {
				nextLoc := [2]int{loc[0] + delta[0], loc[1] + delta[1]}
				if (nextLoc == p.entry) || (nextLoc == p.exit) || (p.inbound(nextLoc) && !nogo[nextLoc]) {
					next[nextLoc] = true
				}
			}
			if !nogo[loc] {
				next[loc] = true
			}
		}
		switch phase {
		case 0:
			if next[p.exit] {
				phase = 1
				next = map[[2]int]bool{p.exit: true}
			}
		case 1:
			if next[p.entry] {
				phase = 2
				next = map[[2]int]bool{p.entry: true}
			}
		case 2:
			if next[p.exit] {
				return t + 1
			}
		}
		possible = next
		t++
	}
}

func init() {
	base.Register(24, &p{})
}

type direction byte

const (
	N direction = iota
	E
	S
	W
)

var deltas = map[direction][2]int{
	N: {-1, 0},
	E: {0, 1},
	S: {1, 0},
	W: {0, -1},
}
