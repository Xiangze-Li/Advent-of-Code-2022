package puzzle07

import (
	"advent2022/base"
	"advent2022/util"
	"fmt"
	"strconv"
	"strings"
)

type p struct {
	lines []string
	root  *dir

	dirSizes map[string]uint64
}

type dir struct {
	name    string
	files   map[string]file
	subdirs map[string]*dir
	parent  *dir
}

func newDir(name string, parent *dir) *dir {
	return &dir{
		name:    name,
		files:   make(map[string]file),
		subdirs: make(map[string]*dir),
		parent:  parent,
	}
}

type file struct {
	name string
	size uint64
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
	p.root = newDir("", nil)
	cwd := p.root

	for i := 0; i < len(p.lines); i++ {
		line := p.lines[i]
		util.Assert(strings.HasPrefix(line, "$"), fmt.Sprintf("line %d: %s", i, line))
		fields := strings.Fields(line)[1:]

		switch fields[0] {
		case "cd":
			switch fields[1] {
			case "/":
				cwd = p.root
			case "..":
				cwd = cwd.parent
			default:
				cwd = cwd.subdirs[fields[1]]
			}
		case "ls":
			for i+1 < len(p.lines) && !strings.HasPrefix(p.lines[i+1], "$") {
				sp := strings.Split(p.lines[i+1], " ")
				if sp[0] == "dir" {
					sub := newDir(sp[1], cwd)
					cwd.subdirs[sp[1]] = sub
				} else {
					cwd.files[sp[1]] = file{
						name: sp[1],
						size: util.Must(strconv.ParseUint(sp[0], 10, 64)),
					}
				}
				i++
			}
		}
	}

	p.dirSizes = make(map[string]uint64)
	p.walkTree("/", p.root, &p.dirSizes)
}

func (p *p) walkTree(cwd string, tree *dir, stat *map[string]uint64) uint64 {
	if tree == nil {
		return 0
	}

	var sum uint64
	for _, f := range tree.files {
		sum += f.size
	}

	for _, sub := range tree.subdirs {
		sum += p.walkTree(cwd+sub.name+"/", sub, stat)
	}

	(*stat)[cwd] = sum
	return sum
}

func (p *p) Solve1() any {
	return util.ReduceMap(p.dirSizes, func(acc uint64, _ string, x uint64) uint64 {
		if x <= 100_000 {
			return acc + x
		}
		return acc
	}, 0)
}

func (p *p) Solve2() any {
	free := 70_000_000 - p.dirSizes["/"]
	toFree := 30_000_000 - free

	return util.ReduceMap(p.dirSizes, func(min uint64, _ string, x uint64) uint64 {
		if toFree <= x && x < min {
			return x
		}
		return min
	}, p.dirSizes["/"])
}

func init() {
	base.Register(7, &p{})
}
